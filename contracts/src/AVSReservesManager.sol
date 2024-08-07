// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import {IPaymentCoordinator} from "@eigenlayer/contracts/interfaces/IPaymentCoordinator.sol";

import {IAVSReservesManager} from "./interfaces/IAVSReservesManager.sol";
import {AVSReservesManagerStorage} from "./AVSReservesManagerStorage.sol";

import {ISafetyFactorOracle} from "./interfaces/ISafetyFactorOracle.sol";
import {IServiceManager} from "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";

import "./static/Structs.sol";
import "./libraries/Accumulator.sol";
import "forge-std/console.sol";

// The AVSReservesManager contract is responsible for managing the token flow to the Service Manager contract
// It is also responsible for updating the token flow based on the Safety Factor
// The Safety Factor is determined by the Safety Factor Oracle contract which represents the protocol's attack surface health

// The reserves manager serves as a 'battery' for the Service Manager contract:
// Storing excess tokens when the protocol is healthy and releasing them when the protocol is in need of more security
contract AVSReservesManager is AVSReservesManagerStorage, AccessControl, Initializable {
    using SafeERC20 for IERC20;
    using AccumulatorLib for Accumulator;

    /**
     *
     *                            Immutables
     *
     */
    IServiceManager public immutable avsServiceManager; // Address of the Service Manager contract

    /**
     *
     *                            Modifiers
     *
     */
    modifier onlyAvsGov() {
        require(hasRole(AVS_GOV_ROLE, msg.sender), "Caller is not a AVS Gov");
        _;
    }

    modifier onlyAnzenGov() {
        require(hasRole(ANZEN_GOV_ROLE, msg.sender), "Caller is not a Anzen Gov");
        _;
    }

    modifier afterEpochExpired() {
        require(
            block.timestamp >= lastEpochUpdateTimestamp + safetyFactorConfig.minEpochDuration, "Epoch not yet expired"
        );
        _;
    }

    constructor(address _avsServiceManager) {
        avsServiceManager = IServiceManager(_avsServiceManager);
    }

    function initialize(
        SafetyFactorConfig memory _safetyFactorConfig,
        address _safetyFactorOracle,
        address _avsGov,
        address _anzenGov,
        uint32 _protocolId,
        address[] memory _rewardTokens,
        uint256[] memory _initial_tokenFlowsPerSecond,
        uint256 _performanceFeeBPS
    ) external initializer {
        _validateSafetyFactorConfig(_safetyFactorConfig);
        // require that the number of reward tokens is equal to the number of initial token flows
        require(_rewardTokens.length == _initial_tokenFlowsPerSecond.length, "Invalid number of reward tokens");

        // require token flows are greater than 0
        for (uint256 index = 0; index < _initial_tokenFlowsPerSecond.length; index++) {
            _validateInitialTokensPerSecond(_initial_tokenFlowsPerSecond[index]);
        }

        safetyFactorConfig = _safetyFactorConfig;

        safetyFactorOracle = ISafetyFactorOracle(_safetyFactorOracle);

        protocolId = _protocolId;
        rewardTokens = _rewardTokens;
        int256 currentSafetyFactor = safetyFactorOracle.getSafetyFactor(protocolId).safetyFactor;

        // initialize token flow for each reward token
        for (uint256 rewardTokenIndex = 0; rewardTokenIndex < _rewardTokens.length; rewardTokenIndex++) {
            rewardTokenAccumulator[_rewardTokens[rewardTokenIndex]].init(
                _initial_tokenFlowsPerSecond[rewardTokenIndex], currentSafetyFactor
            );
        }

        lastEpochUpdateTimestamp = block.timestamp;
        lastPaymentTimestamp = uint32(block.timestamp);
        performanceFeeBPS = _performanceFeeBPS;

        anzen = _anzenGov;

        _grantRole(AVS_GOV_ROLE, _avsGov);
        _grantRole(DEFAULT_ADMIN_ROLE, _avsGov);
        _grantRole(ANZEN_GOV_ROLE, _anzenGov);
    }

    function updateFlow() public afterEpochExpired {
        // This function programmatically adjusts the token flow based on the Safety Factor
        int256 currentSafetyFactor = safetyFactorOracle.getSafetyFactor(protocolId).safetyFactor;

        for (uint256 index = 0; index < rewardTokens.length; index++) {
            rewardTokenAccumulator[rewardTokens[index]].adjustEpochFlow(
                safetyFactorConfig, currentSafetyFactor, performanceFeeBPS, lastEpochUpdateTimestamp
            );
        }

        lastEpochUpdateTimestamp = uint32(block.timestamp);
    }

    // Function to transfer tokenFlow to the Service Manager contract
    function makeRangePaymentsToServiceManager() public {
        IPaymentCoordinator.RangePayment[] memory rangePayments = _createAllRangePayments();

        // Transfer the tokens to the Service Manager contract
        avsServiceManager.payForRange(rangePayments);

        // Update the last payment timestamp
        lastPaymentTimestamp = uint32(block.timestamp);
        lastEpochUpdateTimestamp = uint32(block.timestamp);
    }

    /**
     *
     *                            AVS Governance Functions
     *
     */
    function overrideTokensPerSecond(address[] memory _tokenAddresses, uint256[] memory _newTokensPerSecond)
        external
        onlyAvsGov
    {
        // require that the number of reward tokens is equal to the number of new token flows
        require(_tokenAddresses.length == _newTokensPerSecond.length, "Invalid number of reward tokens");

        // This function is only callable by the AVS delegated address and should only be used in emergency situations
        for (uint256 index = 0; index < _tokenAddresses.length; index++) {
            require(rewardTokenAccumulator[_tokenAddresses[index]].tokensPerSecond != 0, "Reward token does not exist");
            rewardTokenAccumulator[_tokenAddresses[index]].overrideTokensPerSecond(
                _newTokensPerSecond[index], lastEpochUpdateTimestamp
            );
        }
        lastEpochUpdateTimestamp = block.timestamp;
    }

    function overrideTokensPerSecondForToken(address _rewardToken, uint256 _newTokensPerSecond) external onlyAvsGov {
        require(rewardTokenAccumulator[_rewardToken].tokensPerSecond != 0, "Reward token does not exist");
        rewardTokenAccumulator[_rewardToken].overrideTokensPerSecond(_newTokensPerSecond, lastEpochUpdateTimestamp);
        lastEpochUpdateTimestamp = block.timestamp;
    }

    function overrideClaimableTokens(address _rewardToken, uint256 _claimableTokens) external onlyAvsGov {
        // This function is only callable by the AVS delegated address and should only be used in emergency situations
        rewardTokenAccumulator[_rewardToken].overrideClaimableTokens(_claimableTokens);
    }

    function addRewardToken(
        address _rewardToken,
        uint256 _initialTokenFlow,
        IPaymentCoordinator.StrategyAndMultiplier[] memory _strategyAndMultipliers
    ) external onlyAvsGov {
        require(rewardTokenAccumulator[_rewardToken].tokensPerSecond == 0, "Reward token already exists");
        _validateInitialTokensPerSecond(_initialTokenFlow);

        rewardTokens.push(_rewardToken);
        rewardTokenAccumulator[_rewardToken].init(
            _initialTokenFlow, safetyFactorOracle.getSafetyFactor(protocolId).safetyFactor
        );
        _setStrategyAndMultipliers(_rewardToken, _strategyAndMultipliers);
    }

    function removeRewardToken(address _rewardToken) external onlyAvsGov {
        delete rewardTokenAccumulator[_rewardToken];
        delete strategyAndMultipliers[_rewardToken];

        for (uint256 index = 0; index < rewardTokens.length; index++) {
            if (rewardTokens[index] == _rewardToken) {
                rewardTokens[index] = rewardTokens[rewardTokens.length - 1];
                rewardTokens.pop();
                break;
            }
        }
    }

    function updateSafetyFactorParams(SafetyFactorConfig memory _newSafetyFactorConfig) external onlyAvsGov {
        _validateSafetyFactorConfig(_newSafetyFactorConfig);

        safetyFactorConfig = _newSafetyFactorConfig;
    }

    function setStrategyAndMultipliers(
        address _rewardToken,
        IPaymentCoordinator.StrategyAndMultiplier[] memory _strategyAndMultipliers
    ) external onlyAvsGov {
        _setStrategyAndMultipliers(_rewardToken, _strategyAndMultipliers);
    }

    /**
     *
     *                            View Functions
     *
     */
    function getSafetyFactorConfig() external view returns (SafetyFactorConfig memory) {
        return safetyFactorConfig;
    }

    /**
     *
     *                            Anzen Governance Functions
     *
     */
    function adjustFeeBps(uint256 _newFeeBps) external onlyAnzenGov {
        _validatePerformanceFee(_newFeeBps);
        performanceFeeBPS = _newFeeBps;
    }

    /**
     *
     *                            Internal Functions
     *
     */
    function _setStrategyAndMultipliers(
        address _rewardToken,
        IPaymentCoordinator.StrategyAndMultiplier[] memory _strategyAndMultipliers
    ) internal {
        // Clear the current storage array
        delete strategyAndMultipliers[_rewardToken];

        // Push each element from the memory array to the storage array
        for (uint256 index = 0; index < _strategyAndMultipliers.length; index++) {
            strategyAndMultipliers[_rewardToken].push(_strategyAndMultipliers[index]);
        }
    }

    function _createAllRangePayments() internal returns (IPaymentCoordinator.RangePayment[] memory) {
        // Create a range payment for each reward token
        IPaymentCoordinator.RangePayment[] memory rangePayments =
            new IPaymentCoordinator.RangePayment[](rewardTokens.length);

        for (uint256 index = 0; index < rewardTokens.length; index++) {
            (uint256 claimableTokens, uint256 claimableFees) =
                rewardTokenAccumulator[rewardTokens[index]].claim(performanceFeeBPS, lastPaymentTimestamp);
            IERC20 rewardToken = IERC20(rewardTokens[index]);

            _transferPerformanceFeeToAnzen(rewardToken, claimableFees);

            rangePayments[index] = _createRangePayment(rewardToken, claimableTokens, claimableFees);
        }

        return rangePayments;
    }

    function _createRangePayment(IERC20 _rewardToken, uint256 _claimableAmount, uint256 _fee)
        internal
        view
        returns (IPaymentCoordinator.RangePayment memory)
    {
        // Get the claimable tokens for the reward token
        require(_claimableAmount + _fee <= _rewardToken.balanceOf(address(this)), "Insufficient balance");

        return IPaymentCoordinator.RangePayment({
            token: _rewardToken,
            amount: _claimableAmount,
            duration: uint32(block.timestamp) - lastPaymentTimestamp,
            strategiesAndMultipliers: strategyAndMultipliers[address(_rewardToken)],
            startTimestamp: lastPaymentTimestamp
        });
    }

    function _transferPerformanceFeeToAnzen(IERC20 _rewardToken, uint256 _fee) internal {
        // Transfer the fee to the Anzen contract
        if (_fee == 0) {
            return;
        }
        _rewardToken.safeTransfer(anzen, _fee);
    }

    /**
     *
     *                            Validation Functions
     *
     */
    function _validateSafetyFactorConfig(SafetyFactorConfig memory _config) internal pure {
        require(0 <= _config.TARGET_SF_LOWER_BOUND, "Invalid lower bound");
        require(_config.TARGET_SF_UPPER_BOUND <= int256(PRECISION), "Invalid upper bound");
        require(_config.TARGET_SF_LOWER_BOUND < _config.TARGET_SF_UPPER_BOUND, "Invalid Safety Factor Config");
        require(_config.MAX_REDUCTION_FACTOR < PRECISION, "Invalid Reduction Factor");
        require(_config.MAX_INCREASE_FACTOR < PRECISION, "Invalid Increase Factor");
    }

    function _validatePerformanceFee(uint256 _fee) internal pure {
        require(_fee <= MAX_PERFORMANCE_FEE_BPS, "Fee cannot be greater than 5%");
    }

    function _validateInitialTokensPerSecond(uint256 _tokensPerSecond) internal pure {
        require(_tokensPerSecond > 0, "Invalid initial token flow");
    }
}

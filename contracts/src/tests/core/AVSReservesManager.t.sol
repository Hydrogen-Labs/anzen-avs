//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "forge-std/console.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IStrategy} from "@eigenlayer/contracts/interfaces/IPaymentCoordinator.sol";
import "../../static/Structs.sol";
import "../../AVSReservesManager.sol";

// Mocks
import "../mocks/MockSafetyFactorOracle.sol";
import "../mocks/ERC20Mock.sol";
// Harnesses
import "../harnesses/AVSReservesManagerHarness.sol";

contract AVSReservesManagerTests is Test {
    AVSReservesManagerHarness public avsReservesManager;
    AVSReservesManagerHarness public avsReservesManagerImplementation;

    MockSafetyFactorOracle public safetyFactorOracle;
    MockSafetyFactorOracle public safetyFactorOracleImplementation;

    SafetyFactorConfig public safetyFactorConfig;

    address public anzenGov;
    address public proxyAdmin;
    address public avsGov;
    address public avsServiceManager;
    uint32 public avsId;

    ERC20Mock[] public rewardTokenMocks;
    address[] public rewardTokens;
    uint256[] public initialTokenFlows;

    uint256 public MAX_REDUCTION_RATE = PRECISION * 30 / 100;
    uint256 public MAX_INCREASE_RATE = PRECISION * 25 / 100;

    function setUp() public {
        safetyFactorOracle = new MockSafetyFactorOracle();

        anzenGov = address(0x1235);
        proxyAdmin = address(0x789);
        avsGov = address(0x456);
        avsId = uint32(782);
        avsServiceManager = address(0x123);

        // a healthy safety factor in between the bounds
        int256 safetyFactorInit = (int256(PRECISION) * 25) / 100;
        safetyFactorOracle.setSafetyFactor(avsId, safetyFactorInit);

        rewardTokenMocks = new ERC20Mock[](2);
        rewardTokenMocks[0] = new ERC20Mock();
        rewardTokenMocks[1] = new ERC20Mock();

        rewardTokens = new address[](2);
        rewardTokens[0] = address(rewardTokenMocks[0]);
        rewardTokens[1] = address(rewardTokenMocks[1]);

        initialTokenFlows = new uint256[](2);
        initialTokenFlows[0] = 100;
        initialTokenFlows[1] = 200;

        safetyFactorConfig = SafetyFactorConfig(
            (int256(PRECISION) * 20) / 100, // 120% of the current value
            (int256(PRECISION) * 30) / 100, // 140% of the current value
            MAX_REDUCTION_RATE, // 95% of the current value
            MAX_INCREASE_RATE, // 10% of the current value
            3 days
        );

        avsReservesManagerImplementation = new AVSReservesManagerHarness(avsServiceManager);

        avsReservesManager = AVSReservesManagerHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(avsReservesManagerImplementation),
                    proxyAdmin,
                    abi.encodeWithSelector(
                        avsReservesManager.initialize.selector,
                        safetyFactorConfig,
                        address(safetyFactorOracle),
                        avsGov,
                        anzenGov,
                        avsId,
                        rewardTokens,
                        initialTokenFlows
                    )
                )
            )
        );

        rewardTokenMocks[0].mint(address(avsReservesManager), 1_000_000);
        rewardTokenMocks[1].mint(address(avsReservesManager), 1_000_000);
    }

    function test_constructor() public virtual {
        SafetyFactorConfig memory newConfig = avsReservesManager.getSafetyFactorConfig();

        assertEq(newConfig.TARGET_SF_LOWER_BOUND, safetyFactorConfig.TARGET_SF_LOWER_BOUND);
        assertEq(newConfig.TARGET_SF_UPPER_BOUND, safetyFactorConfig.TARGET_SF_UPPER_BOUND);
        assertEq(newConfig.MAX_REDUCTION_FACTOR, safetyFactorConfig.MAX_REDUCTION_FACTOR);
        assertEq(newConfig.MAX_INCREASE_FACTOR, safetyFactorConfig.MAX_INCREASE_FACTOR);
        assertEq(newConfig.minEpochDuration, safetyFactorConfig.minEpochDuration);
        assertEq(avsReservesManager.lastEpochUpdateTimestamp(), 1);

        // Loop through each reward token and perform assertions
        for (uint256 i = 0; i < rewardTokens.length; i++) {
            (
                uint256 claimableTokens,
                uint256 claimableFees,
                uint256 tokensPerSecond,
                uint256 prevTokensPerSecond,
                int256 lastSafetyFactor
            ) = avsReservesManager.rewardTokenAccumulator(rewardTokens[i]);

            // Assert that the values match the expected initial token flows and default values
            assertEq(tokensPerSecond, initialTokenFlows[i]);
            assertEq(prevTokensPerSecond, initialTokenFlows[i]);
            assertEq(claimableTokens, 0);
            assertEq(claimableFees, 0);
            assertEq(lastSafetyFactor, (int256(PRECISION) * 25) / 100);
        }
    }

    function test_updateFlowNoChangeToSafetyFactor(uint256 timeElapsed) public {
        vm.assume(timeElapsed >= 3 days);
        vm.assume(timeElapsed < 180 days);

        vm.warp(timeElapsed + 1); // 1 second default start time

        avsReservesManager.updateFlow();

        // Loop through each reward token and perform assertions
        for (uint256 i = 0; i < rewardTokens.length; i++) {
            (
                uint256 claimableTokens,
                uint256 claimableFees,
                uint256 tokensPerSecond,
                uint256 prevTokensPerSecond,
                int256 safetyFactor
            ) = avsReservesManager.rewardTokenAccumulator(rewardTokens[i]);

            // Assert that the values match the expected initial token flows and default values
            assertEq(tokensPerSecond, initialTokenFlows[i]);
            assertEq(prevTokensPerSecond, initialTokenFlows[i]);
            assertEq(claimableFees, 0);
            assertEq(claimableTokens, timeElapsed * initialTokenFlows[i]);
            assertEq(safetyFactor, (int256(PRECISION) * 25) / 100);
        }
    }

    function test_rejectUpdateBeforeEpoch(uint256 timeElapsed) public {
        vm.assume(timeElapsed < 3 days);

        vm.warp(timeElapsed + 1); // 1 second default start time

        vm.expectRevert("Epoch not yet expired");
        avsReservesManager.updateFlow();
    }

    function test_updateSafetyFactorParams() public {
        SafetyFactorConfig memory newConfig = SafetyFactorConfig(
            (int256(PRECISION) * 10) / 100,
            (int256(PRECISION) * 50) / 100,
            (PRECISION * 90) / 100,
            (PRECISION * 11) / 100,
            4 days
        );

        vm.prank(avsGov);
        avsReservesManager.updateSafetyFactorParams(newConfig);

        SafetyFactorConfig memory updatedConfig = avsReservesManager.getSafetyFactorConfig();

        assertEq(updatedConfig.TARGET_SF_LOWER_BOUND, newConfig.TARGET_SF_LOWER_BOUND);
        assertEq(updatedConfig.TARGET_SF_UPPER_BOUND, newConfig.TARGET_SF_UPPER_BOUND);
        assertEq(updatedConfig.MAX_REDUCTION_FACTOR, newConfig.MAX_REDUCTION_FACTOR);
        assertEq(updatedConfig.MAX_INCREASE_FACTOR, newConfig.MAX_INCREASE_FACTOR);
        assertEq(updatedConfig.minEpochDuration, newConfig.minEpochDuration);
    }

    function test_updateFlowHitsRateLimitReduction() public {
        vm.warp(3 days + 1); // 1 second default start time

        safetyFactorOracle.setSafetyFactor(avsId, int256(PRECISION));
        avsReservesManager.updateFlow();

        // SF is maximal so we should see a reduction in the rate, at the rate limit
        for (uint256 i = 0; i < rewardTokens.length; i++) {
            (,, uint256 tokensPerSecond, uint256 prevTokensPerSecond, int256 safetyFactor) =
                avsReservesManager.rewardTokenAccumulator(rewardTokens[i]);

            // Assert that the values match the expected initial token flows and default values
            assertEq(safetyFactor, int256(PRECISION));
            assertEq(tokensPerSecond, initialTokenFlows[i] * (PRECISION - MAX_REDUCTION_RATE) / PRECISION);
            assertEq(prevTokensPerSecond, initialTokenFlows[i]);
        }
    }

    function test_updateFlowHitsRateLimitIncrease() public {
        vm.warp(3 days + 1); // 1 second default start time

        safetyFactorOracle.setSafetyFactor(avsId, -int256(PRECISION));
        avsReservesManager.updateFlow();

        // SF is minimal so we should see an increase in the rate, at the rate limit
        for (uint256 i = 0; i < rewardTokens.length; i++) {
            (,, uint256 tokensPerSecond, uint256 prevTokensPerSecond, int256 safetyFactor) =
                avsReservesManager.rewardTokenAccumulator(rewardTokens[i]);

            // Assert that the values match the expected initial token flows and default values
            assertEq(safetyFactor, -int256(PRECISION));
            assertEq(tokensPerSecond, initialTokenFlows[i] * (PRECISION + MAX_INCREASE_RATE) / PRECISION);
            assertEq(prevTokensPerSecond, initialTokenFlows[i]);
        }
    }

    function test_updateFlowUnderRateLimitDecrease(int8 sf) public {
        vm.warp(3 days + 1); // 1 second default start time
        vm.assume(sf > 30 && sf < 39);
        int256 newSafetyFactor = (int256(PRECISION) * sf) / 100;
        safetyFactorOracle.setSafetyFactor(avsId, newSafetyFactor);
        avsReservesManager.updateFlow();

        int256 reductionFactor = newSafetyFactor - safetyFactorConfig.TARGET_SF_UPPER_BOUND;

        // SF is in the middle so we should see no change in the rate
        for (uint256 i = 0; i < rewardTokens.length; i++) {
            (,, uint256 tokensPerSecond, uint256 prevTokensPerSecond, int256 safetyFactor) =
                avsReservesManager.rewardTokenAccumulator(rewardTokens[i]);
            uint256 expectedReduction =
                (initialTokenFlows[i] * uint256(reductionFactor)) / uint256(safetyFactorConfig.TARGET_SF_UPPER_BOUND);
            // Assert that the values match the expected initial token flows and default values
            assertEq(safetyFactor, newSafetyFactor);
            assertEq(tokensPerSecond, initialTokenFlows[i] - expectedReduction);
            assertEq(prevTokensPerSecond, initialTokenFlows[i]);
        }
    }

    function test_updateFlowUnderRateLimitIncrease(int8 sf) public {
        vm.warp(3 days + 1); // 1 second default start time
        vm.assume(sf > 15 && sf < 20);
        int256 newSafetyFactor = (int256(PRECISION) * sf) / 100;
        safetyFactorOracle.setSafetyFactor(avsId, newSafetyFactor);
        avsReservesManager.updateFlow();

        int256 increaseFactor = safetyFactorConfig.TARGET_SF_LOWER_BOUND - newSafetyFactor;

        // SF is in the middle so we should see no change in the rate
        for (uint256 i = 0; i < rewardTokens.length; i++) {
            (,, uint256 tokensPerSecond, uint256 prevTokensPerSecond, int256 safetyFactor) =
                avsReservesManager.rewardTokenAccumulator(rewardTokens[i]);
            uint256 expectedIncrease =
                (initialTokenFlows[i] * uint256(increaseFactor)) / uint256(safetyFactorConfig.TARGET_SF_LOWER_BOUND);
            // Assert that the values match the expected initial token flows and default values
            assertEq(safetyFactor, newSafetyFactor);
            assertEq(tokensPerSecond, initialTokenFlows[i] + expectedIncrease);
            assertEq(prevTokensPerSecond, initialTokenFlows[i]);
        }
    }

    function test_overrideTokensPerSecond(uint256 newTokensPerSecond) public {
        address[] memory rewardTokensUpdated = new address[](1);
        rewardTokensUpdated[0] = address(rewardTokens[0]);

        uint256[] memory tokenFlowsUpdated = new uint256[](1);
        tokenFlowsUpdated[0] = newTokensPerSecond;

        vm.prank(avsGov);
        avsReservesManager.overrideTokensPerSecond(rewardTokensUpdated, tokenFlowsUpdated);

        (,, uint256 tokensPerSecond, uint256 prevTokensPerSecond,) =
            avsReservesManager.rewardTokenAccumulator(rewardTokensUpdated[0]);

        assertEq(tokensPerSecond, newTokensPerSecond);
        assertEq(prevTokensPerSecond, newTokensPerSecond);
    }

    function test_overrideTokensPerSecondRevertsNotGov() public {
        address[] memory rewardTokensUpdated = new address[](1);
        rewardTokensUpdated[0] = address(rewardTokens[0]);

        uint256[] memory tokenFlowsUpdated = new uint256[](1);
        tokenFlowsUpdated[0] = 0;

        vm.expectRevert();
        avsReservesManager.overrideTokensPerSecond(rewardTokensUpdated, tokenFlowsUpdated);
    }

    function test_removeRewardToken() public {
        vm.prank(avsGov);
        avsReservesManager.removeRewardToken(rewardTokens[0]);
    }

    function test_removeRewardTokenRevert() public {
        vm.expectRevert();
        avsReservesManager.removeRewardToken(rewardTokens[0]);
    }

    function test_createAllRangePaymentsAllEmpty() public {
        IPaymentCoordinator.RangePayment[] memory rangePayments = avsReservesManager.createAllRangePayments();

        for (uint256 i = 0; i < rangePayments.length; i++) {
            assertEq(rangePayments[i].amount, 0);
        }
    }

    function test_createAllRangePaymentsTimeElapsed(uint256 timeElapsed) public {
        vm.assume(timeElapsed <= 90 days); // safe to assume that the time elapsed between operator payments is less than 90 days
        vm.warp(timeElapsed + 1); // 1 second default start time

        rewardTokenMocks[0].mint(address(avsReservesManager), timeElapsed * initialTokenFlows[0]);
        rewardTokenMocks[1].mint(address(avsReservesManager), timeElapsed * initialTokenFlows[1]);

        IPaymentCoordinator.RangePayment[] memory rangePayments = avsReservesManager.createAllRangePayments();

        for (uint256 i = 0; i < rangePayments.length; i++) {
            assertEq(rangePayments[i].amount, timeElapsed * initialTokenFlows[i]);
            assertEq(rangePayments[i].startTimestamp, 1);
            assertEq(rangePayments[i].duration, timeElapsed);
        }
    }

    function test_setStrategyAndMultipliers() public {
        address rewardToken = address(rewardTokenMocks[0]);

        IPaymentCoordinator.StrategyAndMultiplier[] memory strategies =
            new IPaymentCoordinator.StrategyAndMultiplier[](2);

        strategies[0] = IPaymentCoordinator.StrategyAndMultiplier({strategy: IStrategy(address(0x123)), multiplier: 1});
        strategies[1] = IPaymentCoordinator.StrategyAndMultiplier({strategy: IStrategy(address(0x456)), multiplier: 2});

        vm.prank(avsGov);
        avsReservesManager.setStrategyAndMultipliers(rewardToken, strategies);
    }

    function test_setStrategyAndMultipliersRevert() public {
        address rewardToken = address(rewardTokenMocks[0]);

        IPaymentCoordinator.StrategyAndMultiplier[] memory strategies =
            new IPaymentCoordinator.StrategyAndMultiplier[](1);

        strategies[0] = IPaymentCoordinator.StrategyAndMultiplier({strategy: IStrategy(address(0x123)), multiplier: 1});

        vm.expectRevert();
        avsReservesManager.setStrategyAndMultipliers(rewardToken, strategies);
    }
}
 
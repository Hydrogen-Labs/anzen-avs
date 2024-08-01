// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

import "./interfaces/ISafetyFactorOracle.sol";
import "./AnzenTaskManager.sol";
import "./AVSReservesManager.sol";
import "./AVSReservesManagerFactoryStorage.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract AVSReservesManagerFactory is AVSReservesManagerFactoryStorage {
    ISafetyFactorOracle public immutable safetyFactorOracle;
    address public immutable anzenGov;

    modifier onlyAnzenGov() {
        require(msg.sender == anzenGov, "Only Anzen Gov can call this function");
        _;
    }

    constructor(address _safetyFactorOracle, address _anzenGov) {
        safetyFactorOracle = ISafetyFactorOracle(_safetyFactorOracle);
        anzenGov = _anzenGov;
    }

    function createAVSReservesManager(
        address proxyAdmin,
        SafetyFactorConfig memory _safetyFactorConfig,
        address _avsGov,
        address _avsServiceManager,
        address[] memory _initial_rewardTokens,
        uint256[] memory _initial_tokenFlowsPerSecond,
        uint256 _performanceFeeBPS
    ) external onlyAnzenGov returns (address, address) {
        require(
            !hasAVSReservesManager[_avsServiceManager],
            "AVSReservesManagerFactory: Only one AVSReservesManager per address"
        );

        AVSReservesManager avsReservesManagerImplementation = new AVSReservesManager(_avsServiceManager);

        AVSReservesManager avsReservesManager = AVSReservesManager(
            address(new TransparentUpgradeableProxy(address(avsReservesManagerImplementation), proxyAdmin, ""))
        );

        safetyFactorOracle.addProtocol(lastAVSReservesManagerId, address(avsReservesManager));

        avsReservesManager.initialize(
            _safetyFactorConfig,
            address(safetyFactorOracle),
            _avsGov,
            anzenGov,
            lastAVSReservesManagerId,
            _initial_rewardTokens,
            _initial_tokenFlowsPerSecond,
            _performanceFeeBPS
        );

        hasAVSReservesManager[_avsServiceManager] = true;
        avsReservesManagers[lastAVSReservesManagerId] = address(avsReservesManager);
        lastAVSReservesManagerId++;

        return (address(avsReservesManager), address(avsReservesManagerImplementation));
    }
}

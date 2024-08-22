//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "forge-std/console.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../static/Structs.sol";
import "../../AVSReservesManager.sol";
import "../../AVSReservesManagerFactory.sol";

// Mocks
import "../mocks/MockSafetyFactorOracle.sol";

contract AVSReservesManagerFactoryTests is Test {
    AVSReservesManagerFactory public avsReservesManagerFactory;

    AVSReservesManager public avsReservesManager;
    AVSReservesManager public avsReservesManagerImplementation;

    MockSafetyFactorOracle public safetyFactorOracle;

    SafetyFactorConfig public safetyFactorConfig;

    address public proxyAdmin;
    address public avsGov;
    address public avsServiceManager;
    uint32 public avsId;

    address[] public rewardTokens;
    uint256[] public initialTokenFlows;

    function setUp() public {
        avsGov = address(0x456);
        proxyAdmin = address(0x789);
        avsId = uint32(782);
        avsServiceManager = address(0x123);

        safetyFactorOracle = new MockSafetyFactorOracle();

        avsReservesManagerFactory = new AVSReservesManagerFactory(address(safetyFactorOracle), avsGov);
    }

    function test_createAVSReservesManagerFromFactory() public virtual {
        rewardTokens = new address[](2);
        rewardTokens[0] = address(0xabc);
        rewardTokens[1] = address(0xdef);

        initialTokenFlows = new uint256[](2);
        initialTokenFlows[0] = 100;
        initialTokenFlows[1] = 200;

        safetyFactorConfig = SafetyFactorConfig(
            (int256(PRECISION) * 20) / 100, // 120% of the current value
            (int256(PRECISION) * 40) / 100, // 140% of the current value
            (PRECISION * 95) / 100, // 95% of the current value
            (PRECISION * 10) / 100, // 105% of the current value
            3 days
        );

        vm.prank(avsGov);
        (address avsReservesManagerAddr, address avsReservesManagerImplementationAddr) = avsReservesManagerFactory
            .createAVSReservesManager(
            proxyAdmin, safetyFactorConfig, avsGov, avsServiceManager, rewardTokens, initialTokenFlows
        );

        avsReservesManager = AVSReservesManager(avsReservesManagerAddr);
        avsReservesManagerImplementation = AVSReservesManager(avsReservesManagerImplementationAddr);

        assertEq(avsReservesManagerFactory.avsReservesManagers(0), address(avsReservesManager));
        assertEq(avsReservesManagerFactory.lastAVSReservesManagerId(), 1);
    }

    function test_reverts_createAVSReservesManager_notAnzenGov() public virtual {
        rewardTokens = new address[](2);
        rewardTokens[0] = address(0xabc);
        rewardTokens[1] = address(0xdef);

        initialTokenFlows = new uint256[](2);
        initialTokenFlows[0] = 100;
        initialTokenFlows[1] = 200;

        safetyFactorConfig = SafetyFactorConfig(
            (int256(PRECISION) * 20) / 100, // 120% of the current value
            (int256(PRECISION) * 40) / 100, // 140% of the current value
            (PRECISION * 95) / 100, // 95% of the current value
            (PRECISION * 5) / 100, // 105% of the current value
            3 days
        );

        vm.expectRevert();
        avsReservesManagerFactory.createAVSReservesManager(
            proxyAdmin, safetyFactorConfig, avsGov, avsServiceManager, rewardTokens, initialTokenFlows
        );
    }

    function test_reverts_duplicate_createAVSReservesManager() public virtual {
        rewardTokens = new address[](2);
        rewardTokens[0] = address(0xabc);
        rewardTokens[1] = address(0xdef);

        initialTokenFlows = new uint256[](2);
        initialTokenFlows[0] = 100;
        initialTokenFlows[1] = 200;

        safetyFactorConfig = SafetyFactorConfig(
            (int256(PRECISION) * 20) / 100, // 120% of the current value
            (int256(PRECISION) * 40) / 100, // 140% of the current value
            (PRECISION * 95) / 100, // 95% of the current value
            (PRECISION * 5) / 100, // 105% of the current value
            3 days
        );

        vm.prank(avsGov);
        (address avsReservesManagerAddr, address avsReservesManagerImplementationAddr) = avsReservesManagerFactory
            .createAVSReservesManager(
            proxyAdmin, safetyFactorConfig, avsGov, avsServiceManager, rewardTokens, initialTokenFlows
        );

        avsReservesManager = AVSReservesManager(avsReservesManagerAddr);
        avsReservesManagerImplementation = AVSReservesManager(avsReservesManagerImplementationAddr);

        vm.expectRevert();

        avsReservesManagerFactory.createAVSReservesManager(
            proxyAdmin, safetyFactorConfig, avsGov, avsServiceManager, rewardTokens, initialTokenFlows
        );
    }
}

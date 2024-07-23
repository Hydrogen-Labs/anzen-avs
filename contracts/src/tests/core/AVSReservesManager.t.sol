//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "forge-std/console.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../static/Structs.sol";
import "../../AVSReservesManager.sol";

// Mocks
import "../mocks/MockSafetyFactorOracle.sol";

contract AVSReservesManagerTests is Test {
    AVSReservesManager public avsReservesManager;
    AVSReservesManager public avsReservesManagerImplementation;

    MockSafetyFactorOracle public safetyFactorOracle;
    MockSafetyFactorOracle public safetyFactorOracleImplementation;

    SafetyFactorConfig public safetyFactorConfig;

    address public proxyAdmin;
    address public avsGov;
    address public avsServiceManager;
    uint32 public avsId;

    address[] public rewardTokens;
    uint256[] public initialTokenFlows;

    function setUp() public {
        safetyFactorOracle = new MockSafetyFactorOracle();

        proxyAdmin = address(0x789);
        avsGov = address(0x456);
        avsId = uint32(782);
        avsServiceManager = address(0x123);

        // a healthy safety factor in between the bounds
        int256 safetyFactorInit = (int256(PRECISION) * 130) / 100;
        safetyFactorOracle.setSafetyFactor(avsId, safetyFactorInit);

        rewardTokens = new address[](2);
        rewardTokens[0] = address(0xabc);
        rewardTokens[1] = address(0xdef);

        initialTokenFlows = new uint256[](2);
        initialTokenFlows[0] = 100;
        initialTokenFlows[1] = 200;

        safetyFactorConfig = SafetyFactorConfig(
            (int256(PRECISION) * 120) / 100, // 120% of the current value
            (int256(PRECISION) * 140) / 100, // 140% of the current value
            (PRECISION * 95) / 100, // 95% of the current value
            (PRECISION * 105) / 10, // 105% of the current value
            3 days
        );

        avsReservesManagerImplementation = new AVSReservesManager(avsServiceManager);

        avsReservesManager = AVSReservesManager(
            address(
                new TransparentUpgradeableProxy(
                    address(avsReservesManagerImplementation),
                    proxyAdmin,
                    abi.encodeWithSelector(
                        avsReservesManager.initialize.selector,
                        safetyFactorConfig,
                        address(safetyFactorOracle),
                        avsGov,
                        avsId,
                        rewardTokens,
                        initialTokenFlows
                    )
                )
            )
        );
    }

    function test_constructor() public virtual {
        SafetyFactorConfig memory newConfig = avsReservesManager.getSafetyFactorConfig();

        assertEq(newConfig.TARGET_SF_LOWER_BOUND, safetyFactorConfig.TARGET_SF_LOWER_BOUND);
        assertEq(newConfig.TARGET_SF_UPPER_BOUND, safetyFactorConfig.TARGET_SF_UPPER_BOUND);
        assertEq(newConfig.REDUCTION_FACTOR, safetyFactorConfig.REDUCTION_FACTOR);
        assertEq(newConfig.INCREASE_FACTOR, safetyFactorConfig.INCREASE_FACTOR);
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
            assertEq(lastSafetyFactor, (int256(PRECISION) * 130) / 100);
        }
    }

    function test_updateFlow(uint256 timeElapsed) public {
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
            assertEq(safetyFactor, (int256(PRECISION) * 130) / 100);
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
            (int256(PRECISION) * 110) / 100,
            (int256(PRECISION) * 150) / 100,
            (PRECISION * 90) / 100,
            (PRECISION * 110) / 10,
            4 days
        );

        vm.prank(avsGov);
        avsReservesManager.updateSafetyFactorParams(newConfig);

        SafetyFactorConfig memory updatedConfig = avsReservesManager.getSafetyFactorConfig();

        assertEq(updatedConfig.TARGET_SF_LOWER_BOUND, newConfig.TARGET_SF_LOWER_BOUND);
        assertEq(updatedConfig.TARGET_SF_UPPER_BOUND, newConfig.TARGET_SF_UPPER_BOUND);
        assertEq(updatedConfig.REDUCTION_FACTOR, newConfig.REDUCTION_FACTOR);
        assertEq(updatedConfig.INCREASE_FACTOR, newConfig.INCREASE_FACTOR);
        assertEq(updatedConfig.minEpochDuration, newConfig.minEpochDuration);
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../../../src/AnzenServiceManager.sol" as anzensm;
import {AnzenTaskManager} from "../../../src/AnzenTaskManager.sol";
import {SafetyFactorOracle} from "../../../src/SafetyFactorOracle.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract AnzenTaskManagerTest is BLSMockAVSDeployer {
    anzensm.AnzenServiceManager sm;
    anzensm.AnzenServiceManager smImplementation;

    AnzenTaskManager tm;
    AnzenTaskManager tmImplementation;

    SafetyFactorOracle sfo;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator = address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator = address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    address owner = address(uint160(uint256(keccak256(abi.encodePacked("owner")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        // First, deploy the SafetyFactorOracle contract.
        sfo = new SafetyFactorOracle();

        tmImplementation =
            new AnzenTaskManager(anzensm.IRegistryCoordinator(address(registryCoordinator)), TASK_RESPONSE_WINDOW_BLOCK);

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = AnzenTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector, pauserRegistry, owner, aggregator, generator, address(sfo)
                    )
                )
            )
        );
    }

    function test_CreateNewOraclePullTask() public {
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator, generator);
        tm.createNewOraclePullTask(0, 2, 0, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }

    function test_UpdateTaskGenerator(address newGenerator) public {
        vm.prank(address(owner));
        tm.updateTaskGenerator(newGenerator);

        assertEq(tm.generator(), newGenerator);
    }

    function test_reverts_UpdateTaskGenerator(address newGenerator) public {
        vm.expectRevert();
        tm.updateTaskGenerator(newGenerator);
    }

    function testUpdateAggregator(address newAggregator) public {
        vm.prank(address(owner));
        tm.updateAggregator(newAggregator);

        assertEq(tm.aggregator(), newAggregator);
    }

    function test_reverts_UpdateAggregator(address newAggregator) public {
        vm.expectRevert();
        tm.updateAggregator(newAggregator);
    }
}

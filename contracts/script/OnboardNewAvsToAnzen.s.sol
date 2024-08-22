// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {AnzenServiceManager, IServiceManager} from "../src/AnzenServiceManager.sol";
import {AVSReservesManagerFactory} from "../src/AVSReservesManagerFactory.sol";
import {SafetyFactorConfig} from "../src/SafetyFactorOracle.sol";
import "../../src/tests/mocks/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/OnboardNewAvsToAnzen.s.sol:OnboardNewAvsToAnzen --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract OnboardNewAvsToAnzen is Script, Utils {
    // DEPLOYMENT CONSTANTS

    bytes32 public salt = keccak256(abi.encodePacked(vm.envString("DEPLOYMENT_SALT")));
    AVSReservesManagerFactory public avsReservesManagerFactory;

    address public avsServiceManagerAddr;

    ProxyAdmin public avsProxyAdmin;
    address public newAvsReservesManager;
    address public newAvsReservesManagerImplementation;
    uint32 public avsId;

    function run() external {
        address anzenCommunityMultisig = msg.sender;

        string memory anzenDeployedContracts = readOutput("anzen_avs_deployment_output");

        avsReservesManagerFactory = AVSReservesManagerFactory(
            stdJson.readAddress(anzenDeployedContracts, ".addresses.avsReservesManagerFactory")
        );

        avsId = avsReservesManagerFactory.lastAVSReservesManagerId();

        vm.startBroadcast();
        _onboardAVSToAnzen(anzenCommunityMultisig);
        vm.stopBroadcast();
    }

    function _onboardAVSToAnzen(address anzenCommunityMultisig) internal {
        // deploy proxy admin for ability to upgrade proxy contracts
        avsProxyAdmin = new ProxyAdmin();

        // hard-coded inputs
        _churnSalt();

        // TODO: make all of this inputs configurable with the environment
        SafetyFactorConfig memory safetyFactorConfig = SafetyFactorConfig(200_000, 300_000, 200_000, 200_000, 1 days);

        (newAvsReservesManager, newAvsReservesManagerImplementation) = avsReservesManagerFactory
            .createAVSReservesManager(
            address(avsProxyAdmin),
            safetyFactorConfig,
            anzenCommunityMultisig,
            avsServiceManagerAddr,
            new address[](0),
            new uint256[](0)
        );

        // write the output
        writeAvsOnboardingOutput(
            address(avsProxyAdmin), newAvsReservesManager, newAvsReservesManagerImplementation, avsId
        );
    }

    function _churnSalt() internal {
        salt = keccak256(abi.encode(salt));
    }
}

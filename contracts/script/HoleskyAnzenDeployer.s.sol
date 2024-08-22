// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {AnzenServiceManager, IServiceManager} from "../src/AnzenServiceManager.sol";
import {AVSReservesManagerFactory} from "../src/AVSReservesManagerFactory.sol";
import {AnzenTaskManager} from "../src/AnzenTaskManager.sol";
import {SafetyFactorOracle, SafetyFactorConfig} from "../src/SafetyFactorOracle.sol";
import {IAnzenTaskManager} from "../src/interfaces/IAnzenTaskManager.sol";
import "../../src/tests/mocks/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/AnzenDeployer.s.sol:AnzenDeployer --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract AnzenDeployer is Script, Utils {
    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;

    address public AGGREGATOR_ADDR = vm.envAddress("INIT_AGGREGATOR_ADDR");
    address public TASK_GENERATOR_ADDR = vm.envAddress("INIT_TASK_GENERATOR_ADDR");
    bytes32 public salt = keccak256(abi.encodePacked(vm.envString("DEPLOYMENT_SALT")));

    // ERC20 and Strategy: we need to deploy this erc20, create a strategy for it, and whitelist this strategy in the strategymanager
    StrategyBaseTVLLimits public erc20Strategy;

    // anzen contracts
    ProxyAdmin public anzenProxyAdmin;
    PauserRegistry public anzenPauserReg;

    regcoord.RegistryCoordinator public registryCoordinator;
    regcoord.IRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    AnzenServiceManager public anzenServiceManager;
    IServiceManager public anzenServiceManagerImplementation;

    AnzenTaskManager public anzenTaskManager;
    IAnzenTaskManager public anzenTaskManagerImplementation;

    SafetyFactorOracle public safetyFactorOracle;
    SafetyFactorOracle public safetyFactorOracleImplementation;

    AVSReservesManagerFactory public avsReservesManagerFactory;
    AVSReservesManagerFactory public avsReservesManagerFactoryImplementation;

    address public anzenReservesManager;
    address public anzenReservesManagerImplementation;

    function run() external {
        address delegationManagerAddr = 0xA44151489861Fe9e3055d95adC98FbD462B948e7;
        address avsDirectoryAddr = 0x055733000064333CaDDbC92763c58BF0192fFeBf;
        address wethStrategyAddr = 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9;

        // Eigenlayer contracts
        IDelegationManager delegationManager = IDelegationManager(delegationManagerAddr);
        IAVSDirectory avsDirectory = IAVSDirectory(avsDirectoryAddr);

        erc20Strategy = StrategyBaseTVLLimits(wethStrategyAddr);

        address anzenCommunityMultisig = msg.sender;
        address anzenPauser = msg.sender;

        vm.startBroadcast();
        _deployAnzenContracts(delegationManager, avsDirectory, erc20Strategy, anzenCommunityMultisig, anzenPauser);
        vm.stopBroadcast();
    }

    function _deployAnzenContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        IStrategy strat,
        address anzenCommunityMultisig,
        address anzenPauser
    ) internal {
        // Adding this as a temporary fix to make the rest of the script work with a single strategy
        // since it was originally written to work with an array of strategies
        IStrategy[1] memory deployedStrategyArray = [strat];
        uint256 numStrategies = deployedStrategyArray.length;

        // deploy proxy admin for ability to upgrade proxy contracts
        anzenProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = anzenPauser;
            pausers[1] = anzenCommunityMultisig;
            anzenPauserReg = new PauserRegistry{salt: salt}(pausers, anzenCommunityMultisig);
        }

        EmptyContract emptyContract = new EmptyContract{salt: salt}();

        // hard-coded inputs

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        anzenServiceManager = AnzenServiceManager(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        _churnSalt();
        anzenTaskManager = AnzenTaskManager(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        _churnSalt();
        registryCoordinator = regcoord.RegistryCoordinator(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        _churnSalt();
        blsApkRegistry = IBLSApkRegistry(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        _churnSalt();
        indexRegistry = IIndexRegistry(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        _churnSalt();
        stakeRegistry = IStakeRegistry(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        _churnSalt();
        safetyFactorOracle = SafetyFactorOracle(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        _churnSalt();
        avsReservesManagerFactory = AVSReservesManagerFactory(
            address(new TransparentUpgradeableProxy{salt: salt}(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        // Current step

        operatorStateRetriever = new OperatorStateRetriever{salt: salt}();

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        {
            stakeRegistryImplementation = new StakeRegistry{salt: salt}(registryCoordinator, delegationManager);

            anzenProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))), address(stakeRegistryImplementation)
            );

            blsApkRegistryImplementation = new BLSApkRegistry{salt: salt}(registryCoordinator);

            anzenProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))), address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry{salt: salt}(registryCoordinator);

            anzenProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(indexRegistry))), address(indexRegistryImplementation)
            );
        }

        registryCoordinatorImplementation = new regcoord.RegistryCoordinator{salt: salt}(
            anzenServiceManager,
            regcoord.IStakeRegistry(address(stakeRegistry)),
            regcoord.IBLSApkRegistry(address(blsApkRegistry)),
            regcoord.IIndexRegistry(address(indexRegistry))
        );

        {
            uint256 numQuorums = 1;
            // for each quorum to setup, we need to define
            // QuorumOperatorSetParam, minimumStakeForQuorum, and strategyParams
            regcoord.IRegistryCoordinator.OperatorSetParam[] memory quorumsOperatorSetParams =
                new regcoord.IRegistryCoordinator.OperatorSetParam[](numQuorums);
            for (uint256 i = 0; i < numQuorums; i++) {
                // hard code these for now
                quorumsOperatorSetParams[i] = regcoord.IRegistryCoordinator.OperatorSetParam({
                    maxOperatorCount: 10000,
                    kickBIPsOfOperatorStake: 15000,
                    kickBIPsOfTotalStake: 100
                });
            }
            // set to 0 for every quorum
            uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
            IStakeRegistry.StrategyParams[][] memory quorumsStrategyParams =
                new IStakeRegistry.StrategyParams[][](numQuorums);
            for (uint256 i = 0; i < numQuorums; i++) {
                quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](numStrategies);
                for (uint256 j = 0; j < numStrategies; j++) {
                    quorumsStrategyParams[i][j] = IStakeRegistry.StrategyParams({
                        strategy: deployedStrategyArray[j],
                        // setting this to 1 ether since the divisor is also 1 ether
                        // therefore this allows an operator to register with even just 1 token
                        // see https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/src/StakeRegistry.sol#L484
                        //    weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
                        multiplier: 1 ether
                    });
                }
            }
            anzenProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(registryCoordinator))),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    regcoord.RegistryCoordinator.initialize.selector,
                    // we set churnApprover and ejector to communityMultisig because we don't need them
                    anzenCommunityMultisig,
                    anzenCommunityMultisig,
                    anzenCommunityMultisig,
                    anzenPauserReg,
                    0, // 0 initialPausedStatus means everything unpaused
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
                )
            );
        }

        anzenServiceManagerImplementation =
            new AnzenServiceManager{salt: salt}(avsDirectory, registryCoordinator, stakeRegistry, anzenTaskManager);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        anzenProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(anzenServiceManager))),
            address(anzenServiceManagerImplementation)
        );

        anzenTaskManagerImplementation =
            new AnzenTaskManager{salt: salt}(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK);

        safetyFactorOracleImplementation = new SafetyFactorOracle{salt: salt}();

        // TODO: Replace 3rd param with fallbackposter address
        anzenProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(safetyFactorOracle))),
            address(safetyFactorOracleImplementation),
            abi.encodeWithSelector(
                safetyFactorOracleImplementation.initialize.selector,
                address(anzenTaskManager),
                address(avsReservesManagerFactory),
                anzenCommunityMultisig
            )
        );

        avsReservesManagerFactoryImplementation =
            new AVSReservesManagerFactory{salt: salt}(address(safetyFactorOracle), anzenCommunityMultisig);

        anzenProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(avsReservesManagerFactory))),
            address(avsReservesManagerFactoryImplementation)
        );

        SafetyFactorConfig memory safetyFactorConfig = SafetyFactorConfig(200_000, 300_000, 200_000, 200_000, 1 days);

        (anzenReservesManager, anzenReservesManagerImplementation) = avsReservesManagerFactory.createAVSReservesManager(
            address(anzenProxyAdmin),
            safetyFactorConfig,
            anzenCommunityMultisig,
            address(anzenServiceManager),
            new address[](0),
            new uint256[](0)
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        anzenProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(anzenTaskManager))),
            address(anzenTaskManagerImplementation),
            abi.encodeWithSelector(
                anzenTaskManager.initialize.selector,
                anzenPauserReg,
                anzenCommunityMultisig,
                AGGREGATOR_ADDR,
                TASK_GENERATOR_ADDR,
                address(safetyFactorOracle)
            )
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "anzenProxyAdmin", address(anzenProxyAdmin));
        vm.serializeAddress(deployed_addresses, "erc20Strategy", address(erc20Strategy));
        vm.serializeAddress(deployed_addresses, "anzenServiceManager", address(anzenServiceManager));
        vm.serializeAddress(
            deployed_addresses, "anzenServiceManagerImplementation", address(anzenServiceManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "anzenTaskManager", address(anzenTaskManager));
        vm.serializeAddress(
            deployed_addresses, "anzenTaskManagerImplementation", address(anzenTaskManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "registryCoordinator", address(registryCoordinator));
        vm.serializeAddress(
            deployed_addresses, "registryCoordinatorImplementation", address(registryCoordinatorImplementation)
        );
        vm.serializeAddress(deployed_addresses, "safetyFactorOracle", address(safetyFactorOracle));
        vm.serializeAddress(
            deployed_addresses, "safetyFactorOracleImplementation", address(safetyFactorOracleImplementation)
        );
        vm.serializeAddress(deployed_addresses, "avsReservesManagerFactory", address(avsReservesManagerFactory));
        vm.serializeAddress(
            deployed_addresses,
            "avsReservesManagerFactoryImplementation",
            address(avsReservesManagerFactoryImplementation)
        );

        string memory deployed_addresses_output =
            vm.serializeAddress(deployed_addresses, "operatorStateRetriever", address(operatorStateRetriever));

        // serialize all the data
        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);

        writeOutput(finalJson, "anzen_avs_deployment_output");
        writeAvsOnboardingOutput(address(anzenProxyAdmin), anzenReservesManager, anzenReservesManagerImplementation, 0);
    }

    function _churnSalt() internal {
        salt = keccak256(abi.encode(salt));
    }
}

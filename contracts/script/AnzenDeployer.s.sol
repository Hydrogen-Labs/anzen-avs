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
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
    address public constant TASK_GENERATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    // ERC20 and Strategy: we need to deploy this erc20, create a strategy for it, and whitelist this strategy in the strategymanager

    ERC20Mock public erc20Mock;
    StrategyBaseTVLLimits public erc20MockStrategy;

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

    address public anzenAvsReservesManager;
    address public anzenAvsReservesManagerImplementation;

    function run() external {
        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput("eigenlayer_deployment_output");
        IStrategyManager strategyManager =
            IStrategyManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.strategyManager"));
        IDelegationManager delegationManager =
            IDelegationManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.delegation"));
        IAVSDirectory avsDirectory =
            IAVSDirectory(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.avsDirectory"));
        ProxyAdmin eigenLayerProxyAdmin =
            ProxyAdmin(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerProxyAdmin"));
        PauserRegistry eigenLayerPauserReg =
            PauserRegistry(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerPauserReg"));
        StrategyBaseTVLLimits baseStrategyImplementation = StrategyBaseTVLLimits(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.baseStrategyImplementation")
        );

        address anzenCommunityMultisig = msg.sender;
        address anzenPauser = msg.sender;

        vm.startBroadcast();
        _deployErc20AndStrategyAndWhitelistStrategy(
            eigenLayerProxyAdmin, eigenLayerPauserReg, baseStrategyImplementation, strategyManager
        );
        _deployAnzenContracts(delegationManager, avsDirectory, erc20MockStrategy, anzenCommunityMultisig, anzenPauser);
        vm.stopBroadcast();
    }

    function _deployErc20AndStrategyAndWhitelistStrategy(
        ProxyAdmin eigenLayerProxyAdmin,
        PauserRegistry eigenLayerPauserReg,
        StrategyBaseTVLLimits baseStrategyImplementation,
        IStrategyManager strategyManager
    ) internal {
        erc20Mock = new ERC20Mock();
        // TODO(samlaf): any reason why we are using the strategybase with tvl limits instead of just using strategybase?
        // the maxPerDeposit and maxDeposits below are just arbitrary values.
        erc20MockStrategy = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        StrategyBaseTVLLimits.initialize.selector,
                        1 ether, // maxPerDeposit
                        100 ether, // maxDeposits
                        IERC20(erc20Mock),
                        eigenLayerPauserReg
                    )
                )
            )
        );
        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = erc20MockStrategy;
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        thirdPartyTransfersForbiddenValues[0] = false;
        strategyManager.addStrategiesToDepositWhitelist(strats, thirdPartyTransfersForbiddenValues);
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
            anzenPauserReg = new PauserRegistry(pausers, anzenCommunityMultisig);
        }

        EmptyContract emptyContract = new EmptyContract();

        // hard-coded inputs

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        anzenServiceManager = AnzenServiceManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );
        anzenTaskManager = AnzenTaskManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );
        registryCoordinator = regcoord.RegistryCoordinator(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );
        blsApkRegistry = IBLSApkRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );
        indexRegistry = IIndexRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );
        stakeRegistry = IStakeRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );
        safetyFactorOracle = SafetyFactorOracle(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );
        avsReservesManagerFactory = AVSReservesManagerFactory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(anzenProxyAdmin), ""))
        );

        operatorStateRetriever = new OperatorStateRetriever();

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        {
            stakeRegistryImplementation = new StakeRegistry(registryCoordinator, delegationManager);

            anzenProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))), address(stakeRegistryImplementation)
            );

            blsApkRegistryImplementation = new BLSApkRegistry(registryCoordinator);

            anzenProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))), address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry(registryCoordinator);

            anzenProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(indexRegistry))), address(indexRegistryImplementation)
            );
        }

        registryCoordinatorImplementation = new regcoord.RegistryCoordinator(
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
            new AnzenServiceManager(avsDirectory, registryCoordinator, stakeRegistry, anzenTaskManager);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        anzenProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(anzenServiceManager))),
            address(anzenServiceManagerImplementation)
        );

        anzenTaskManagerImplementation = new AnzenTaskManager(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK);

        safetyFactorOracleImplementation = new SafetyFactorOracle();

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
            new AVSReservesManagerFactory(address(safetyFactorOracle), anzenCommunityMultisig);

        anzenProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(avsReservesManagerFactory))),
            address(avsReservesManagerFactoryImplementation)
        );

        SafetyFactorConfig memory safetyFactorConfig = SafetyFactorConfig(200_000, 300_000, 200_000, 200_000, 1 days);

        (anzenAvsReservesManager, anzenAvsReservesManagerImplementation) = avsReservesManagerFactory
            .createAVSReservesManager(
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
        vm.serializeAddress(deployed_addresses, "erc20Mock", address(erc20Mock));
        vm.serializeAddress(deployed_addresses, "erc20MockStrategy", address(erc20MockStrategy));
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

        writeAvsOnboardingOutput(
            address(anzenProxyAdmin), anzenAvsReservesManager, anzenAvsReservesManagerImplementation, 0
        );
    }
}

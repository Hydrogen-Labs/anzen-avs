# Anzen AVS

<b> Do not use it in Production, testnet only. </b>

## Dependencies

You will need [foundry](https://book.getfoundry.sh/getting-started/installation) and [zap-pretty](https://github.com/maoueh/zap-pretty) and docker to run the examples below.

```
curl -L https://foundry.paradigm.xyz | bash
foundryup
go install github.com/maoueh/zap-pretty@latest
```

You will also need to [install docker](https://docs.docker.com/get-docker/), and build the contracts:

```
make build-contracts
```

## Running locally via make

This simple session illustrates the basic flow of the AVS. The makefile commands are hardcoded for a single operator, but it's however easy to create new operator config files, and start more operators manually (see the actual commands that the makefile calls).

Start anvil in a separate terminal:

```bash
make start-anvil-chain-with-el-and-avs-deployed
```

The above command starts a local anvil chain from a [saved state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json) with eigenlayer and anzen contracts already deployed (but no operator registered).

Start the aggregator:

```bash
make start-aggregator
```

Register the operator with eigenlayer and the anzen avs, and then start the process:

```bash
make start-operator
```

> By default, the `start-operator` command will also setup the operator (see `register_operator_on_startup` flag in `config-files/anvil/operator.yaml`). To disable this, set `register_operator_on_startup` to false, and run `make cli-setup-operator` before running `start-operator`.

## Running via docker compose

We wrote a [docker-compose.yml](./docker-compose.yml) file to run and test everything on a single machine. It will start an anvil instance, loading a [state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json) where the eigenlayer and anzen contracts are deployed, start the aggregator, and finally one operator, along with prometheus and grafana servers. The grafana server will be available at http://localhost:3000, with user and password both set to `admin`. We have created a simple [grafana dashboard](./grafana/provisioning/dashboards/AVSs/anzen.json) which can be used as a starting example and expanded to include AVS specific metrics. The eigen metrics should not be added to this dashboard as they will be exposed on the main eigenlayer dashboard provided by the eigenlayer-cli.

## Running on Holesky testnet

Refer to the [Holesky Guide](./docs/holesky-guide.md) for instructions on how to run the AVS on the Holesky testnet.

## Avs Task Description

The architecture of the AVS contains:

- [Eigenlayer core](https://github.com/Layr-Labs/eigenlayer-contracts/tree/master) contracts
- AVS contracts
  - [ServiceManager](contracts/src/AnzenServiceManager.sol) which will eventually contain slashing logic but for M2 is just a placeholder.
  - [TaskManager](contracts/src/AnzenTaskManager.sol) which contains [task creation](contracts/src/AnzenTaskManager.sol#L83) and [task response](contracts/src/AnzenTaskManager.sol#L102) logic.
  - The [challenge](contracts/src/AnzenTaskManager.sol#L176) logic could be separated into its own contract, but we have decided to include it in the TaskManager for this simple task.
  - Set of [registry contracts](https://github.com/Layr-Labs/eigenlayer-middleware) to manage operators opted in to this avs
- Task Generator
  - in a real world scenario, this could be a separate entity, but for this simple demo, the aggregator also acts as the task generator
- Aggregator
  - aggregates BLS signatures from operators and posts the aggregated response to the task manager
  - For this simple demo, the aggregator is not an operator, and thus does not need to register with eigenlayer or the AVS contract. It's IP address is simply hardcoded into the operators' config.
- Operators
  - verify the proposed safety factor response, sign it, and send it to the aggregator

1. A task generator (in our case, same as the aggregator) publishes tasks once every regular interval (say 10 blocks, you are free to set your own interval) to the AnzenTaskManager contract's [createNewTask](contracts/src/AnzenTaskManager.sol#L83) function. Each task specifies a protocol `oracleIndex` for which it wants the currently opted-in operators to determine its safety factor `safetyFactor`. `createNewTask` also takes `quorumNumbers` and `quorumThresholdPercentage` which requests that each listed quorum (we only use quorumNumber 0 in anzen) needs to reach at least thresholdPercentage of operator signatures.

2. A [registry](https://github.com/Layr-Labs/eigenlayer-middleware/blob/master/src/BLSRegistryCoordinatorWithIndices.sol) contract is deployed that allows any eigenlayer operator with at least 1 delegated [mockerc20](contracts/src/ERC20Mock.sol) token to opt-in to this AVS and also de-register from this AVS.

3. [Operator] The operators who are currently opted-in with the AVS need to read the task number from the Task contract, compute its safety factor, sign on that computed result (over the BN254 curve) and send their taskResponse and signature to the aggregator.

4. [Aggregator] The aggregator collects the signatures from the operators and aggregates them using BLS aggregation. If any response passes the [quorumThresholdPercentage](contracts/src/interfaces/IAnzenTaskManager.sol#L36) set by the task generator when posting the task, the aggregator posts the aggregated response to the Task contract.

5. If a response was sent within the [response window](contracts/src/AnzenTaskManager.sol#L119), we enter the [Dispute resolution] period.
   - [Off-chain] A challenge window is launched during which anyone can [raise a dispute](contracts/src/AnzenTaskManager.sol#L171) in a DisputeResolution contract (in our case, this is the same as the TaskManager contract)
   - [On-chain] The DisputeResolution contract resolves that a particular operator’s response is not the correct response (that is, not the square of the integer specified in the task) or the opted-in operator didn’t respond during the response window. If the dispute is resolved, the operator will be frozen in the Registration contract and the veto committee will decide whether to veto the freezing request or not.

Below is a more detailed uml diagram of the aggregator and operator processes:

## Avs node spec compliance

Every AVS node implementation is required to abide by the [Eigenlayer AVS Node Specification](https://docs.eigenlayer.xyz/category/node-specification). We suggest reading through the whole spec, including the keys management section, but the hard requirements are currently only to:

- implement the [AVS Node API](https://docs.eigenlayer.xyz/category/avs-node-api)
- implement the [eigen prometheus metrics](https://docs.eigenlayer.xyz/category/metrics)

If you are using golang, you can use our [metrics](https://github.com/Layr-Labs/eigensdk-go/tree/master/metrics) and [nodeapi](https://github.com/Layr-Labs/eigensdk-go/tree/master/nodeapi) implementation in the [eigensdk](https://github.com/Layr-Labs/eigensdk-go), just like this repo does. Otherwise, you will have to implement it on your own.

## StakeUpdates Cronjob

AVS Registry contracts have a stale view of operator shares in the delegation manager contract. In order to update their stake table, they need to periodically call the [StakeRegistry.updateStakes()](https://github.com/Layr-Labs/eigenlayer-middleware/blob/f171a0812126bbb0bb6d44f53c622591a643e987/src/StakeRegistry.sol#L76) function. We are currently writing a cronjob binary to do this for you, will be open sourced soon!

## Integration Tests

See the integration tests [README](tests/anvil/README.md) for more details.

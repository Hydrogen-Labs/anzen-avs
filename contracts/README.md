## Smart Contract Architecture

```mermaid
graph TD

    F[Anzen Operators] -->|Oracle Consensus| A[Anzen Task Manager]
    G[Anzen Aggregator] -->|Initializes Tasks| A
    A -->|Operator Consensus| B((Safety Factor Oracle))
    B --> |Adjusts Emissions| D1((AVS1 Reserves))
    B --> |Adjusts Emissions| D2((AVS2 Reserves))
    B --> |Adjusts Emissions| D3((AVS3 Reserves))
    C[AVS Reserves Factory] --> |Creates Reserves Manager| D1
    C --> |Creates Reserves Manager| D2
    C --> |Creates Reserves Manager| D3
    C --> |Register AVS| B

    subgraph AVS1 Controlled Contracts
        D1 --> |Sends Payment| E1((AVS1 Reward Manager))
    end

    subgraph AVS2 Controlled Contracts
        D2 --> |Sends Payment| E2((AVS2 Reward Manager))
    end

    subgraph AVS3 Controlled Contracts
        D3 --> |Sends Payment| E3((AVS3 Reward Manager))
    end
```

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```

### Anvil

```shell
$ anvil
```

### Deploy

```shell
$ make deploy-anzen-contracts-to-anvil-and-save-state
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```

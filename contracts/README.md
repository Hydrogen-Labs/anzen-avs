## Architecture

```mermaid
graph TD
    F[Anzen Operators] -->|Oracle Consensus| A[Anzen Task Manager]
    A -->|Operator Consensus| B((Safety Factor Oracle))
    B --> |Adjusts Emissions| D((AVS Reserves))
    C[AVS Reserves Factory] --> |Creates Reserves Manager| D
    C --> |Register AVS| B

    subgraph AVS Controlled Contracts
        D --> |Sends Payment| E((AVS Reward Manager))
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
$ forge script script/Counter.s.sol:CounterScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Cast

```shell
$ cast <subcommand>
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```

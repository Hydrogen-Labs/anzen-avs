## Troubleshooting

### Received error from aggregator

When running on anvil, a typical log for the operator is

```
[2024-04-09 18:25:08.647 PDT] INFO (logging/zap_logger.go:49) rpc client is nil. Dialing aggregator rpc client
[2024-04-09 18:25:08.650 PDT] INFO (logging/zap_logger.go:49) Sending signed task response header to aggregator {"signedTaskResponse":"\u0026aggregator.SignedTaskResponse{TaskResponse:contractAnzenTaskManager.IAnzenTaskManagerTaskResponse{ReferenceTaskIndex:0x2, NumberSquared:4}, BlsSignature:bls.Signature{G1Point:(*bls.G1Point)(0x14000282068)}, OperatorId:[32]uint8{0xc4, 0xc2, 0x10, 0x30, 0xe, 0x28, 0xab, 0x4b, 0xa7, 0xb, 0x7f, 0xbb, 0xe, 0xfa, 0x55, 0x7d, 0x2a, 0x2a, 0x5f, 0x1f, 0xbf, 0xa6, 0xf8, 0x56, 0xe4, 0xcf, 0x3e, 0x9d, 0x76, 0x6a, 0x21, 0xdc}}"}
[2024-04-09 18:25:08.651 PDT] INFO (logging/zap_logger.go:49) Received error from aggregator {"err":"task 2 not initialized or already completed"}
[2024-04-09 18:25:08.651 PDT] INFO (logging/zap_logger.go:69) Retrying in 2 seconds
[2024-04-09 18:25:10.679 PDT] INFO (logging/zap_logger.go:49) Signed task response header accepted by aggregator. {"reply":false}
```

The error `task 2 not initialized or already completed` is expected behavior. This is because the aggregator needs to setup its data structures before it can accept responses. But on a local anvil setup, the operator had time to receive the websocket event for the new task, square the number, sign the response, and send it to the aggregator process before the aggregator has finalized its setup. Hence, the operator retries sending the response 2 seconds later and it is accepted.

For bindings and abigen if you have not already done so, you will need to install the go tools:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Issue with forge: `error failed to read artifact source when running`

try rebuilding the contracts:

```bash
cd contracts
forge clean
forge build
```

### Make `tests-integration` failing

If running on an a mac:

https://github.com/foundry-rs/foundry/issues/8039

```bash
docker pull --platform linux/amd64 ghcr.io/foundry-rs/foundry:latest
```

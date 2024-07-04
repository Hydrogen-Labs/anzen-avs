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

# Anzen Holesky Guide

<b> Do not use it in Production, testnet only. </b>

## Joining the Holesky AVS

Review the guide here: [Operator Guide](../docs/operator-guide.md)

## Deploying the Holesky contracts

### Create the deployment `.env` file

```bash
cd contracts
cp sample.env .env
```

Edit the .env file with your variables

### Build and deploy the contracts

```bash
make build-contracts
make deploy-anzen-contracts-to-holesky-and-save-state
```

This will deploy the contracts to the holesky chain and save the state to [`contracts/script/output/17000/holesky_anzen_avs_deployment_output.json`](../contracts/script/output/17000/holesky_anzen_avs_deployment_output.json)

## Running the Holesky AVS

### Create operator keys

Follow the Eigenlayer guide to create keys [here](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-keys), for holesky deployment it is okay to use the `--insecure` flag in the key creation command. You will need both an `edcsa` and `bls` key.

### Create .env file

```bash
cd config-files/holesky
cp sample.aggregator.yaml aggregator.yaml
cp sample.operator.yaml operator.yaml
```

Edit the aggregator.yaml and operator.yaml files with your variables:

[aggregator.yaml](../config-files/holesky/aggregator.yaml)

[operator.yaml](../config-files/holesky/operator.yaml)

### Start the Holesky AVS

```bash
make start-aggregator-holesky
```

in a separate terminal

```bash
make start-operator-holesky
```

## Troubleshooting

#### RPC Issues

Ensure that your RPC client let's you create websocket subscriptions.
Recommended RPC providers are:

- [Tenderly](https://tenderly.co/)

The free tier of Tenderly is sufficient for testing.

#### Operator not joining the AVS

- Ensure that you have sufficient eth in your operator account to pay for gas costs.
- Ensure that you have wETH in your operator account to pay for joining the AVS (this is for the default strategy of the AVS)

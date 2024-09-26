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

This will deploy the contracts to the holesky chain and save the state to [`contracts/script/output/17000/anzen_avs_deployment_output.json`](../contracts/script/output/17000/anzen_avs_deployment_output.json)

## Running the Holesky AVS

### Create operator keys

Follow the Eigenlayer guide to create keys [here](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-keys), for holesky deployment it is okay to use the `--insecure` flag in the key creation command. You will need both an `edcsa` and `bls` key.

### Create config files

You only need to set up the operator.yaml file for the operator. If you want to set up an aggregator you will need to set up the aggregator.yaml file as well, but this is not required to join the AVS.

```bash
cd config-files/holesky
cp sample.operator.yaml operator.yaml
```

Optional:

```bash
cp sample.aggregator.yaml aggregator.yaml
```

Edit the aggregator.yaml and operator.yaml files with your variables:

[aggregator.yaml](../config-files/holesky/aggregator.yaml)

[operator.yaml](../config-files/holesky/operator.yaml)

### Start the Holesky AVS using Docker

#### Operator

in a separate terminal

```bash
docker compose -f operator.docker-compose.yml up
```

#### Aggregator

Configure the aggregator environment variables with the aggregator private key

```bash
cp sample.aggregator.env .env
```

```bash
docker compose -f aggregator.docker-compose.yml up
```

### Start the Holesky AVS using Make

```bash
make start-operator-holesky
```

Optional:

```bash
make start-aggregator-holesky
```

## Troubleshooting

First, you should check for any updates to the operator image

```bash
docker pull ghcr.io/hydrogen-labs/anzen-avs/aggregator/cmd/main.go:latest
docker pull ghcr.io/hydrogen-labs/anzen-avs/operator/cmd/main.go:latest
```

#### RPC Issues

Ensure that your RPC client let's you create websocket subscriptions.
Recommended RPC providers are:

- [Tenderly](https://tenderly.co/)

The free tier of Tenderly is sufficient for testing.

#### Operator not joining the AVS

- Ensure that you have sufficient eth in your operator account to pay for gas costs.
- Ensure that you have wETH in your operator account to pay for joining the AVS (this is for the default strategy of the AVS)

# This imports the environment variables from the .env file
include .env
export $(shell sed 's/=.*//' .env)
############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

AGGREGATOR_ECDSA_PRIV_KEY=0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
CHALLENGER_ECDSA_PRIV_KEY=0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a

CHAINID=31337
HOLESKY_CHAINID=17000
# Make sure to update this if the strategy address changes
# check in contracts/script/output/${CHAINID}/anzen_avs_deployment_output.json
STRATEGY_ADDRESS=0x7a2088a1bFc9d81c55368AE168C2C02570cB814F
DEPLOYMENT_FILES_DIR=contracts/script/output/${CHAINID}
HOLESKY_FILES_DIR=contracts/script/output/${HOLESKY_CHAINID}

-----------------------------: ## 

___CONTRACTS___: ## 

build-contracts: ## builds all contracts
	cd contracts && forge build

deploy-eigenlayer-contracts-to-anvil-and-save-state: ## Deploy eigenlayer
	./tests/anvil/deploy-eigenlayer-save-anvil-state.sh

deploy-anzen-contracts-to-anvil-and-save-state: ## Deploy avs
	./tests/anvil/deploy-avs-save-anvil-state.sh

deploy-anzen-contracts-to-holesky-and-save-state: ## Deploy avs
	./scripts/deploy-anzen-holesky.sh

deploy-all-to-anvil-and-save-state: deploy-eigenlayer-contracts-to-anvil-and-save-state deploy-anzen-contracts-to-anvil-and-save-state ## deploy eigenlayer, shared avs contracts, and inc-sq contracts 

start-anvil-chain-with-el-and-avs-deployed: ## starts anvil from a saved state file (with el and avs contracts deployed)
	./tests/anvil/start-anvil-chain-with-el-and-avs-deployed.sh

onboard-avs-to-anvil-anzen:
	./tests/anvil/onboard-avs-to-anzen.sh

bindings: ## generates contract bindings
	cd contracts && ./generate-go-bindings.sh

___DOCKER___: ## 
docker-build-and-publish-images: ## builds and publishes operator and aggregator docker images using Ko
	KO_DOCKER_REPO=ghcr.io/hydrogen-labs/anzen-avs ko build aggregator/cmd/main.go --preserve-import-paths
	KO_DOCKER_REPO=ghcr.io/hydrogen-labs/anzen-avs ko build operator/cmd/main.go --preserve-import-paths
docker-compose-up: ## runs docker compose pull first and then up
	docker compose pull && docker compose up -d

__CLI__: ## 

cli-setup-operator: send-fund cli-register-operator-with-eigenlayer cli-deposit-into-mocktoken-strategy cli-register-operator-with-avs ## registers operator with eigenlayer and avs

cli-register-operator-with-eigenlayer: ## registers operator with delegationManager
	go run cli/main.go --config config-files/anvil/operator.yaml register-operator-with-eigenlayer

cli-deposit-into-mocktoken-strategy: ## 
	./scripts/deposit-into-mocktoken-strategy.sh

cli-register-operator-with-avs: ## 
	go run cli/main.go --config config-files/anvil/operator.yaml register-operator-with-avs

cli-deregister-operator-with-avs: ## 
	go run cli/main.go --config config-files/anvil/operator.yaml deregister-operator-with-avs

cli-print-operator-status: ## 
	go run cli/main.go --config config-files/anvil/operator.yaml print-operator-status

send-fund: ## sends fund to the operator saved in tests/keys/test.ecdsa.key.json
	cast send 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5 --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

-----------------------------: ## 
# We pipe all zapper logs through https://github.com/maoueh/zap-pretty so make sure to install it
# TODO: piping to zap-pretty only works when zapper environment is set to production, unsure why
____OFFCHAIN_SOFTWARE___: ## 
start-aggregator: ## 
	go run aggregator/cmd/main.go --config config-files/anvil/aggregator.yaml \
		--anzen-deployment ${DEPLOYMENT_FILES_DIR}/anzen_avs_deployment_output.json \
		--ecdsa-private-key ${AGGREGATOR_ECDSA_PRIV_KEY} \
		2>&1 | zap-pretty

start-aggregator-holesky:
	go run aggregator/cmd/main.go --config config-files/holesky/aggregator.yaml \
			--anzen-deployment ${HOLESKY_FILES_DIR}/anzen_avs_deployment_output.json \
			--ecdsa-private-key ${ECDSA_PRIVATE_KEY} \
			2>&1 | zap-pretty


start-operator: ## 
	go run operator/cmd/main.go --config config-files/anvil/operator.yaml \
		--anzen-deployment ${DEPLOYMENT_FILES_DIR}/anzen_avs_deployment_output.json \
		2>&1 | zap-pretty

start-operator-holesky: ## 
	go run operator/cmd/main.go --config config-files/holesky/operator.yaml \
		--anzen-deployment ${HOLESKY_FILES_DIR}/anzen_avs_deployment_output.json \
		2>&1 | zap-pretty

start-operator-2: ## 
	go run operator/cmd/main.go --config config-files/anvil/operator-2.yaml \
		--anzen-deployment ${DEPLOYMENT_FILES_DIR}/anzen_avs_deployment_output.json \
		2>&1 | zap-pretty

start-operator-3: ## 
	go run operator/cmd/main.go --config config-files/anvil/operator-3.yaml \
		--anzen-deployment ${DEPLOYMENT_FILES_DIR}/anzen_avs_deployment_output.json \
		2>&1 | zap-pretty

start-challenger: ## 
	go run challenger/cmd/main.go --config config-files/anvil/challenger.yaml \
		--anzen-deployment ${DEPLOYMENT_FILES_DIR}/anzen_avs_deployment_output.json \
		--ecdsa-private-key ${CHALLENGER_ECDSA_PRIV_KEY} \
		2>&1 | zap-pretty

run-plugin: ## 
	go run plugin/cmd/main.go --config config-files/anvil/operator.yaml
-----------------------------: ## 
_____HELPER_____: ## 
mocks: ## generates mocks for tests
	go install go.uber.org/mock/mockgen@v0.3.0
	go generate ./...

tests-unit: ## runs all unit tests
	go test $$(go list ./... | grep -v /integration) -coverprofile=coverage.out -covermode=atomic --timeout 15s
	go tool cover -html=coverage.out -o coverage.html

tests-contract: ## runs all forge tests
	cd contracts && forge test

tests-integration: ## runs all integration tests
	go test ./tests/integration/... -v -count=1

tests-all: tests-unit tests-contract tests-integration ## runs all tests


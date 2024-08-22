cd contracts
source .env
forge script script/OnboardNewAvsToAnzen.s.sol --rpc-url $RPC_URL --private-key $DEPLOYER_PRIVATE_KEY --broadcast --sender $DEPLOYER_ADDR
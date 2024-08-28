cd contracts

#!/bin/bash

# Load the .env file
set -a
source .env
set +a

# Extract the numeric part of DEPLOYMENT_SALT and increment it
CURRENT_SALT=$(echo $DEPLOYMENT_SALT | grep -o '[0-9]*')
NEW_SALT=$((CURRENT_SALT + 1))

# Update DEPLOYMENT_SALT in .env file (macOS version)
sed -i '' "s/DEPLOYMENT_SALT=\"Deployment$CURRENT_SALT\"/DEPLOYMENT_SALT=\"Deployment$NEW_SALT\"/" .env

# Run the forge command
forge script script/HoleskyAnzenDeployer.s.sol --rpc-url $RPC_URL --private-key $DEPLOYER_PRIVATE_KEY --broadcast --sender $DEPLOYER_ADDR

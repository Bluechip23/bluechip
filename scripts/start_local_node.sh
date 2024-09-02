#!/bin/bash

# debug
# set -x

source ./common.sh
HOME_DIR=$HOME/.bluechip

update_test_genesis () {
  cat $HOME_DIR/config/genesis.json | jq "$1" > $HOME_DIR/config/tmp_genesis.json && mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json
}

# Check if screen installed
check_screen_installed

# Check if any of the named screens already exist and prompt to stop them
prompt_to_stop_screens genesis

# Delete home folder
sudo rm -rf $HOME_DIR

# Define the paths
BLUE_CHIP_REPO=$HOME/bluechip/

# Get the current working directory
current_dir=$(pwd)

if [ ! -d "$BLUE_CHIP_REPO" ]; then
  # Take action if $DIR exists. #
  echo "Need to clone bluechip repo at ${BLUE_CHIP_REPO}"
  exit 1
fi

echo "Initializing genesis node..."

cd $BLUE_CHIP_REPO && COSMOS_BUILD_OPTIONS=nostrip && make install
bluechipd init --chain-id bluechip-1 validator1 --home $HOME_DIR
echo "pet apart myth reflect stuff force attract taste caught fit exact ice slide sheriff state since unusual gaze practice course mesh magnet ozone purchase" | bluechipd keys add validator1 --keyring-backend test --recover --home $HOME_DIR
echo "bottom soccer blue sniff use improve rough use amateur senior transfer quarter" | bluechipd keys add validator2 --keyring-backend test --recover --home $HOME_DIR

bluechipd add-genesis-account validator1 100000000000000ubluechip --keyring-backend test --home $HOME_DIR
bluechipd add-genesis-account validator2 110000000000000ubluechip  --keyring-backend test --home $HOME_DIR
bluechipd gentx validator1 90000000000000ubluechip --keyring-backend test --chain-id bluechip-1
bluechipd collect-gentxs --home $HOME_DIR
sed -i 's/stake/ubluechip/g' ~/.bluechip/config/genesis.json
sed -i 's/pruning-interval = "0"/pruning-interval = "100"/g' $HOME/.bluechip/config/app.toml
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0ubluechip"/g' $HOME/.bluechip/config/app.toml
sed -i 's/127.0.0.1:26657/0.0.0.0:26657/g' $HOME/.bluechip/config/config.toml
sed -i 's/cors_allowed_origins\s*=\s*\[\]/cors_allowed_origins = ["*",]/g' $HOME/.bluechip/config/config.toml
sed -i 's/prometheus = false/prometheus = true/g' $HOME/.bluechip/config/config.toml
sed -i 's/timeout_commit = "5s"/timeout_commit = "4s"/g' $HOME/.bluechip/config/config.toml

bluechipd config keyring-backend test --home "$HOME/.bluechip" >/dev/null 2>&1
bluechipd config broadcast-mode sync --home "$HOME/.bluechip" >/dev/null 2>&1

echo "Updating genesis paramsters."

# Update gov params
update_test_genesis '.app_state["slashing"]["params"]["min_signed_per_window"]="0.050000000000000000"'
update_test_genesis '.app_state["slashing"]["params"]["slash_fraction_double_sign"]="0.080000000000000000"'

echo "Starting genesis node in a screen session..."

screen -dmS genesis bluechipd start --rpc.laddr "tcp://0.0.0.0:26657"

# sleep 5s to wait bluechip L1 RPC gets ready
echo "Bluechip L1 network started. sleep 5s to wait bluechip blockchain rpc gets ready"

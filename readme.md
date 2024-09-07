# Bluechip
# Become a Validator
This assumes that you have go and a proper operating system installed to your machine. If you need to install go, do so here: https://go.dev/doc/install

Once everything is installed, you will need to build the BlueChip binary:
# from $HOME dir
git clone https://github.com/Bluechip23/bluechip

cd bluechip

git fetch

git checkout <version-tag>

Once you find the correct tag, you can then complete installation.

# Make you are in the bluechip directory

make install

# Confirm installation:

bluechipd version

It is important to have a sufficient machine to run a validating node.

# The minimum recommended hardware requirements for running a validator for the BlueChip mainnet are:

4 Cores (modern CPU's)

32GB RAM

1TB of storage (SSD or NVME)

BlueChip is a young and growing chain. As more smart contracts and history develop, these requirements may need upgrading.

# Set the required mainnet chain-id

CHAIN_ID=bluechip-2


# Please decide what you would like your validator to be named:

MONIKER_NAME=<moniker-name>

# Initialize chain:

bluechipd init "$MONIKER_NAME" --chain-id $CHAIN_ID

# This will generate the following files in ~/.bluechip/config/

genesis.json

node_key.json

priv_validator_key.json

# Once you have these files, you will then need to download the genesis.json file:

rm ~/.bluechip/config/genesis.json

wget https://raw.githubusercontent.com/Bluechip23/bluechip/main/genesis.json

mv bluechip-genesis.json $HOME/.bluechip/config/genesis.json


By running this you will replace the genesis file you previously created with the one belonging to the mainnet.

# Find Peers For Network
Find Peers from the BlueChip network and place them in your config file to allow proper node communication.

CHAIN_REPO="https://raw.githubusercontent.com/BlueChip23/bluechip/main" && \
export PEERS="$(curl -sL "$CHAIN_REPO/peers.txt")"

# Set peers in your newly created config file
sed -i.bak -e "s/^peers *=.*/peers = \"$PEERS\"/" ~/.bluechip/config/config.toml

# If you are a validator node, please set a minimum gas price in your app.toml file. Gas fees are paid in blue chips:

sed -i.bak -e "s/^minimum-gas-prices *=.*/minimum-gas-prices = \"0.0025ubluechip,\ ~/.bluechip/config/app.toml

# Once your basic node parameters are set, you can now create a new key pair or restore one.

bluechipd keys add <key-name>

or

bluechipd keys add <key-name> --recover

*Note: you will be prompted to input your mnemonic seed phrase

# Query the keystore for your public address

bluechipd keys show <key-name> -a

# You can name your key a specific name by replacing <key-name> with your chosen name

*NOTE: please write your seed phrase down and store it securely. It will be the only way you can access your keys in the future.

# Now keys are set up, you are able to obtain blue chips, have them bonded to your node, and begin validating. Either find people who are willing to delegate to you, or you will need to purchase blue chips on the open market.

Once you have blue chips you are then able to begin syncing to the chain via the genesis file. 

After running the bluechipd daemon, the chain will begin to sync to the network. The time to sync to the network will vary depending on your setup and the current size of the blockchain. 

To query the status of your node:

curl http://localhost:26657/status | jq .result.sync_info.catching_up

If the following command returns true then your node is still catching up. If it returns false then your node has caught up to the network current block and you are safe to proceed to upgrade to a validator node.

# Once your node is caught up, you are able upgrade your node to a validator.

bluechipd tx staking create-validator 

  --amount 1000000ubluechip 
  
  --commission-max-change-rate "0.1" 
  
  --commission-max-rate "0.20" 
  
  --commission-rate "0.1" 
  
  --min-self-delegation "1" 
  
  --details "validators write bios too" 
  
  --pubkey=$(bluechipd tendermint show-validator) 
  
  --moniker "$MONIKER_NAME" 
  
  --chain-id $CHAIN_ID 
  
  --gas-prices 0.025ubluechip 
  
  --from <key-name>

If you would like to supply more information you can run the command below for more flags: 
bluechipd tx staking create-validator --help


It is best practice to backup important files in case something happens to your validator node. Please do so for the best experience for yourself, delegators, and our chain.

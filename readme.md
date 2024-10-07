# Bluechip
To install and run the BlueChip binary the following guide is applicable. Follow the steps and create your node.

## Install dependencies
```
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install -y build-essential curl wget jq make gcc chrony git
sudo su -c "echo 'fs.file-max = 65536' >> /etc/sysctl.conf"
sudo sysctl -p
```

## Install go
```
sudo rm -rf /usr/local/.go
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
sudo cp /usr/local/go /usr/local/.go -r
sudo rm -rf /usr/local/go
```

## Update environment variables to include go
```
cat <<'EOF' >>$HOME/.profile
export GOROOT=/usr/local/.go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/.go/bin:$HOME/go/bin
EOF
source $HOME/.profile

go version
```
The last command should return something like go version go1.16.4 linux/amd64.

## Install the binary
```
git clone https://github.com/Bluechip23/bluechip.git
cd bluechip
git fetch
git checkout
make install

bluechipd version
```
The last command should return the bluechip binary version.

## Configure the node
Add an address:
```
bluechipd keys add <key-name> 
```

Add the chain-id:
```
bluechipd config chain-id bluechip-2
```

Set the name of your node:
```
bluechipd init <your_custom_moniker> --chain-id bluechip-2
```

Download the genesis file:
```
curl https://raw.githubusercontent.com/Bluechip23/bluechip/main/genesis.json > ~/.bluechipd/config/genesis.json
```

Set the peers in your config.toml. Find peers for the network and place them in your config file to allow proper node communication.
```
CHAIN_REPO="https://raw.githubusercontent.com/BlueChip23/bluechip/main" && \
export PEERS="$(curl -sL "$CHAIN_REPO/peers.txt")"
```

Set peers in your newly created config file
```
sed -i.bak -e "s/^peers *=.*/peers = \"$PEERS\"/" ~/.bluechip/config/config.toml
```

Set a minimum gas price in your app.toml file. Gas fees are paid in blue chips:
```
sed -i.bak -e "s/^minimum-gas-prices *=.*/minimum-gas-prices = \"0.0025ubluechip,\ ~/.bluechip/config/app.toml
```

## Create a service file
```
sudo tee /etc/systemd/system/bluechipd.service > /dev/null <<EOF  
[Unit]
Description=Bluechip Daemon
After=network-online.target

[Service]
User=$USER
ExecStart=$(which bluechipd) start
Restart=always
RestartSec=3
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
EOF

bluechipd tendermint unsafe-reset-all --home $HOME/.bluechipd --keep-addr-book

sudo mv /etc/systemd/system/bluechipd.service /lib/systemd/system/
sudo -S systemctl daemon-reload
sudo -S systemctl enable bluechipd
```

## Start the binary
```
sudo service bluechipd start
```

Monitor the logs of the binary:
```
sudo journalctl -u bluechipd -f
```

To see if your node is caught up:
```
curl http://localhost:26657/status | jq .result.sync_info.catching_up
```
If the following command returns true then your node is still catching up. If it returns false then your node has caught up to the network current block and you are safe to proceed to upgrade to a validator node.

## Start the validator node
If your node is caught up, you can create a validator node:
```
bluechipd tx staking create-validator 
  --amount 1000000ubluechip 
  --commission-max-change-rate "0.1" 
  --commission-max-rate "0.20" 
  --commission-rate "0.1" 
  --min-self-delegation "1" 
  --details "validators write bios too" 
  --pubkey=$(bluechipd tendermint show-validator) 
  --moniker "$MONIKER_NAME" 
  --chain-id bluechip-2
  --gas-prices 0.025ubluechip 
  --from <key-name>
```

It is best practice to backup important files in case something happens to your validator node. Please do so for the best experience for yourself, delegators, and our chain. Do this with the relevant validator files.

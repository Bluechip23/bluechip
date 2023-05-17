# 1. Install the prerequistics.

```
sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
```

## Install Golang:

## Install latest go version https://golang.org/doc/install
```
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.18
source ~/.profile
```

## to verify that Golang installed
```
go version
export PATH=$PATH:/usr/local/go/bin
```

# 2.Install the executables

## clone the bluechip git repository.

```
git clone https://github.com/Smartdev0328/bluechip.git
```

## build the executable binary

```
sudo rm -rf ~/.bluechip
make install
export PATH=$PATH:/home/ubuntu/go/bin
```

## initialize the chain.

```
bluechipd init --chain-id=bluechip_1 validator1
```

## add genesis accounts - backend accout and first validator account
<!-- bluechipd keys add validator1 -->
<!-- echo "seed words for validator1" | bluechipd keys add validator1 --keyring-backend test --recover -->

```
echo "claim either tribe mercy genre drastic stamp spring attend ready believe material hedgehog space remind valley give slight cram arm release universe hybrid abuse" | bluechipd keys add validator1 --keyring-backend test --recover

bluechipd add-genesis-account $(bluechipd keys show validator1 -a --keyring-backend test) 50000000ubluechip 
```
<!-- bluechipd keys add backend_account -->
<!-- echo "seed words for backend wallet" | bluechipd keys add backend_account --keyring-backend test --recover -->
echo "require resist steak energy armed prison embody abuse huge submit host subway merit kiwi inherit distance cliff suffer general program connect link employ crew" | bluechipd keys add backend_account --keyring-backend test --recover

bluechipd add-genesis-account $(bluechipd keys show backend_account -a --keyring-backend test) 30000000000000ubluechip

bluechipd gentx validator1 50000000ubluechip --keyring-backend test --chain-id bluechip_1

bluechipd collect-gentxs
sed -i 's/stake/ubluechip/g' ~/.bluechip/config/genesis.json

## Config parameters
cd ~/.bluechip/config
jq '.app_state.slashing.params.min_signed_per_window = "0.050000000000000000"' genesis.json > temp.json && mv temp.json genesis.json
jq '.app_state.slashing.params.slash_fraction_double_sign = "0.080000000000000000"' genesis.json > temp.json && mv temp.json genesis.json

# 3. start the chain
bluechipd start
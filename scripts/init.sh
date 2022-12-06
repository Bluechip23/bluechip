# bluechipchain deployment script

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
```

## Install the executables


sudo rm -rf ~/.bluechip
make install

bluechipd init --chain-id=bluechip_1 validator1

#bluechipd keys add validator1
#echo "claim either tribe mercy genre drastic stamp spring attend ready believe material hedgehog space remind valley give slight cram arm release universe hybrid abuse" | bluechipd keys add validator1 --keyring-backend os --recover

bluechipd add-genesis-account $(bluechipd keys show validator1 -a --keyring-backend os) 10000000000000ubluechip 

#bluechipd keys add account2
#echo "require resist steak energy armed prison embody abuse huge submit host subway merit kiwi inherit distance cliff suffer general program connect link employ crew" | bluechipd keys add account1 --keyring-backend os --recover

bluechipd add-genesis-account $(bluechipd keys show account1 -a --keyring-backend os) 20000000000000ubluechip

bluechipd gentx validator1 50000000ubluechip --keyring-backend os --chain-id bluechip_1

bluechipd collect-gentxs

sed -i 's/stake/ubluechip/g' ~/.bluechip/config/genesis.json

# cd ~/.bluechip/config
# jq '.app_state.slashing.params.min_signed_per_window = "0.050000000000000000"' genesis.json > temp.json && mv temp.json genesis.json
# jq '.app_state.slashing.params.slash_fraction_double_sign = "0.080000000000000000"' genesis.json > temp.json && mv temp.json genesis.json

bluechipd start
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
bluechipd init --chain-id=bluechip_1 <moniker>
cp genesis.json ~/.bluechip/config
```

## Config parameters
Add seed nodes or persistant peers on config.toml.
persistent_peers="<first_node>"

# 3. start the chain
bluechipd start
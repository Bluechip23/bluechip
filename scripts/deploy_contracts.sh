#!/bin/bash

# Run this from the root repo directory

## CONFIG
# NOTE: you will need to update these to deploy on different network
BINARY='bluechipd'
DENOM='ubluechip'
CHAIN_ID='bluechip_1'
RPC='http://localhost:26657/'
REST='http://localhost:1317/'
TXFLAG="--gas-prices 0.025$DENOM --gas auto --gas-adjustment 1.3 -y -b block --chain-id $CHAIN_ID --node $RPC"
QFLAG="--chain-id $CHAIN_ID --node $RPC"
A="bluechip1wgeeup5y8v04qny99rk7l3d6ygrn9n402eh3q3"
echo "TX Flags: $TXFLAG"

(echo "siren window salt bullet cream letter huge satoshi fade shiver permit offer happy immense wage fitness goose usual aim hammer clap about super trend") | $BINARY keys add test --recover

#### CW20-GOV ####
# Upload cw20 contract code
echo xxxxxxxxx | $BINARY tx wasm store "artifacts/cw20_base.wasm" --from validator1 $TXFLAG
CW20_CODE=$($BINARY q wasm list-code --reverse --output json $QFLAG | jq -r '.code_infos[0].code_id')
echo $CW20_CODE


# Instantiate cw20 contract
CW20_INIT='{
    "name": "PageToken1",
    "symbol": "PGT1",
    "decimals": 6,
    "initial_balances": [{"address":"'"$A"'","amount":"1000000000"},{"address":"'"$($BINARY keys show validator1 -a)"'","amount":"1000000000"}]
}'
echo "$CW20_INIT"
echo xxxxxxxxx | $BINARY tx wasm instantiate $CW20_CODE "$CW20_INIT" --from "validator1" --label "token1" $TXFLAG --no-admin

# Get cw20 contract address
CW20_CONTRACT=$($BINARY q wasm list-contract-by-code $CW20_CODE --output json $QFLAG | jq -r '.contracts[-1]')
echo $CW20_CONTRACT

# Upload cw-dao contract code
echo xxxxxxxxx | $BINARY tx wasm store "artifacts/wasmswap.wasm" --from validator1 $TXFLAG
WASMSWAP_CODE=$($BINARY q wasm list-code --reverse --output json $QFLAG | jq -r '.code_infos[0].code_id')
echo $WASMSWAP_CODE

# Initialize factory contract
SWAP_1_INIT='{
    "token1_denom": {"native": "'$DENOM'"},
    "token2_denom": {"cw20": "'"$CW20_CONTRACT"'"},
    "lp_token_code_id": '$CW20_CODE',
    "owner": "'$A'",
    "protocol_fee_recipient": "'$A'",
    "protocol_fee_percent": "0.2",
    "lp_fee_percent": "0.3"
}'

echo "$SWAP_1_INIT"
echo xxxxxxxxx | $BINARY tx wasm instantiate $WASMSWAP_CODE "$SWAP_1_INIT" --from "validator1" --label "swap_1" $TXFLAG --no-admin
SWAP_1_CONTRACT=$($BINARY q wasm list-contract-by-code $WASMSWAP_CODE --output json $QFLAG | jq -r '.contracts[-1]')
echo $SWAP_1_CONTRACT

# $BINARY tx wasm execute $CW20_CONTRACT '{"increase_allowance":{"amount":"100000000","spender":"'"$SWAP_1_CONTRACT"'"}}' --from test $TXFLAG
# $BINARY tx wasm execute $SWAP_1_CONTRACT '{"add_liquidity":{"token1_amount":"100000000","max_token2":"100000000","min_liquidity":"1"}}' --from test --amount "100000000"$DENOM $TXFLAG

# Instantiate cw20 contract
CW20_INIT_2='{
    "name": "PageToken2",
    "symbol": "PGT2",
    "decimals": 6,
    "initial_balances": [{"address":"'"$A"'","amount":"1000000000"}]
}'
echo "$CW20_INIT_2"
echo xxxxxxxxx | $BINARY tx wasm instantiate $CW20_CODE "$CW20_INIT_2" --from "validator1" --label "token2" $TXFLAG --no-admin

# Get cw20 contract address
CW20_CONTRACT_2=$($BINARY q wasm list-contract-by-code $CW20_CODE --output json $QFLAG | jq -r '.contracts[-1]')
echo $CW20_CONTRACT_2

# Initialize factory contract
SWAP_2_INIT='{
    "token1_denom": {"native": "'$DENOM'"},
    "token2_denom": {"cw20": "'"$CW20_CONTRACT_2"'"},
    "lp_token_code_id": '$CW20_CODE',
    "owner": "'$A'",
    "protocol_fee_recipient": "'$A'",
    "protocol_fee_percent": "0.2",
    "lp_fee_percent": "0.3"
}'
echo "$SWAP_2_INIT"
echo xxxxxxxxx | $BINARY tx wasm instantiate $WASMSWAP_CODE "$SWAP_2_INIT" --from "validator1" --label "swap_2" $TXFLAG --no-admin

SWAP_2_CONTRACT=$($BINARY q wasm list-contract-by-code $WASMSWAP_CODE --output json $QFLAG | jq -r '.contracts[-1]')
# $BINARY tx wasm execute $CW20_CONTRACT_2 '{"increase_allowance":{"amount":"100000000","spender":"'"$SWAP_2_CONTRACT"'"}}' --from test $TXFLAG
# $BINARY tx wasm execute $SWAP_2_CONTRACT '{"add_liquidity":{"token1_amount":"100000000","max_token2":"100000000","min_liquidity":"1"}}' --from test --amount "100000000"$DENOM $TXFLAG

#!/bin/bash

CHAIN_ID="eurx-1"
KEY_PASSPHRASE=''
VALIDATOR_NAME=""
VALIDATOR_ADDRESS=""
VALIDATOR_MNEMONIC=""

GENESIS_FILE_NAME="genesis.json"
GENESIS_TMP_FILE_NAME="$GENESIS_FILE_NAME.tmp"
DENOM="uestm"
GENESIS_AMOUNT="500000000000$DENOM"

function add_key() {
  expect -c "
  set timeout 5
  spawn docker run -v $HOME/.eurxd:/root/.eurxd -v $HOME/.eurxcli:/root/.eurxcli -it eurx eurxcli keys add $1 --recover
  expect \"> Enter your bip39 mnemonic\"
  send \"$2\n\"
  expect \"Enter keyring passphrase:\"
  send \"$3\n\"
  expect \"Re-enter keyring passphrase:\"
  send \"$3\n\"
  interact
  "
}

function gen_tx() {
  sudo mkdir -p ~/.eurxd/config/gentx
  expect -c "
  set timeout 5
  spawn docker run -v $HOME/.eurxd:/root/.eurxd -v $HOME/.eurxcli:/root/.eurxcli -it eurx eurxd gentx --amount $1 --name $2 --output-document "$3"
  expect \"Enter keyring passphrase:\"
  send \"$4\n\"
  expect \"Enter keyring passphrase:\"
  send \"$4\n\"
  expect \"Enter keyring passphrase:\"
  send \"$4\n\"
  interact
  "
}

sudo rm -rf ~/.eurxd ~/.eurxcli

docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxd init eurx --chain-id "$CHAIN_ID"
docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxcli config chain-id "$CHAIN_ID"
docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxcli config trust-node true

add_key "$VALIDATOR_NAME" "$VALIDATOR_MNEMONIC" "$KEY_PASSPHRASE"

docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxd add-genesis-account $VALIDATOR_ADDRESS "$GENESIS_AMOUNT"

gen_tx "$GENESIS_AMOUNT" "$VALIDATOR_NAME" "/root/.eurxd/config/gentx/gentx-validator.json"  "$KEY_PASSPHRASE"
docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxd collect-gentxs

sudo cp ~/.eurxd/config/genesis.json "./$GENESIS_FILE_NAME"
sudo chmod 777 "./$GENESIS_FILE_NAME"

# Update genesis.json
jq ".app_state.mint.params.mint_denom |= \"$DENOM\"" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.staking.params.bond_denom |= \"$DENOM\"" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.crisis.constant_fee.denom |= \"$DENOM\"" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.gov.deposit_params.min_deposit[0].denom |= \"$DENOM\"" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.cdp.gov_denom |= \"$DENOM\"" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.cdp.params.collateral_params += [{\"auction_size\":\"50000000000\",\"conversion_factor\":\"8\",\"debt_limit\":{\"amount\":\"2000000000000\",\"denom\":\"eurx\"},\"denom\":\"bnb\",\"liquidation_penalty\":\"0.05\",\"liquidation_ratio\":\"1.5\",\"spot_market_id\":\"bnb:eur\",\"liquidation_market_id\":\"bnb:eur:30\",\"prefix\":1,\"stability_fee\":\"1.0000000007829977\"}]" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.cdp.params.global_debt_limit.amount |= \"2000000000000\"" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.incentive.params.active |= true" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.incentive.params.rewards += [{\"active\":true,\"denom\":\"bnb\",\"available_rewards\":{\"amount\":\"50000000000\",\"denom\":\"$DENOM\"},\"duration\":\"36288000000000000\",\"time_lock\":\"1892160000000000000\",\"claim_duration\":\"36288000000000000\"}]" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
jq ".app_state.pricefeed.params.markets += [{\"active\":true,\"base_asset\":\"bnb\",\"market_id\":\"bnb:eur\",\"oracles\":[],\"quote_asset\":\"eur\"},{\"active\":true,\"base_asset\":\"bnb\",\"market_id\":\"bnb:eur:30\",\"oracles\":[],\"quote_asset\":\"eur\"}]" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME
# Add validator address to oracles
jq ".app_state.pricefeed.params.markets[].oracles += [\"$VALIDATOR_ADDRESS\"]" $GENESIS_FILE_NAME > $GENESIS_TMP_FILE_NAME && mv $GENESIS_TMP_FILE_NAME $GENESIS_FILE_NAME

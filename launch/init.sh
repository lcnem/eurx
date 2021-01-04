#!/bin/bash

CHAIN_ID="eurx-1"
KEY_PASSPHRASE=''
VALIDATOR_NAME=""
VALIDATOR_ADDRESS=""
VALIDATOR_MNEMONIC=""

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

docker build -t eurx ../

docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxd init eurx --chain-id "$CHAIN_ID"
docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxcli config chain-id "$CHAIN_ID"
docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxcli config trust-node true
add_key "$VALIDATOR_NAME" "$VALIDATOR_MNEMONIC" "$KEY_PASSPHRASE"

KEY_EXISTS=$(jq ".app_state.pricefeed.params.markets[0].oracles | contains([\"$VALIDATOR_ADDRESS\"])" ./genesis.json)
if [ $KEY_EXISTS = "false" ]; then
  jq ".app_state.pricefeed.params.markets[].oracles += [\"$VALIDATOR_ADDRESS\"]" ./genesis.json > ./genesis.json.tmp
  sudo mv ./genesis.json.tmp ~/.eurxd/config/genesis.json
else
  sudo cp ./genesis.json ~/.eurxd/config/genesis.json
fi

docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxd add-genesis-account $VALIDATOR_ADDRESS "500000000000ujsmn,500000000000token"
gen_tx "500000000000ujsmn" "$VALIDATOR_NAME" "/root/.eurxd/config/gentx/gentx-validator.json"  "$KEY_PASSPHRASE"
docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli -it eurx eurxd collect-gentxs

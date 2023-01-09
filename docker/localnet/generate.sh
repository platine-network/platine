#!/bin/bash

# Generates network configuration for an arbitrary 
# amount of validators, observers, and seeds.

set -euo pipefail

# sed in macos requires extra argument
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    SED_EXT=''
elif [[ "$OSTYPE" == "darwin"* ]]; then
    SED_EXT='.orig'
fi

# Params
CHAIN_ID=${1:-"platine-1"} # First parameter, default is "platine-1"

VALIDATORS_COUNT=${2:-4} # Second parameter, default is 4
SEEDS_COUNT=${3:-1} # Third parameter, default is 1
OBSERVERS_COUNT=${4:-1} # Fourth parameter, default is 1

function init_node() {
  NODE_HOME=$1
  NODE_MONIKER=$2

  echo "[${NODE_MONIKER}] Initializing"

  platined init "${NODE_MONIKER}" --chain-id "${CHAIN_ID}" --home "${NODE_HOME}" 2> /dev/null
  platined tendermint show-node-id --home "${NODE_HOME}" > "${NODE_HOME}/node_id.txt"
  platined tendermint show-validator --home "${NODE_HOME}" > "${NODE_HOME}/node_val_pubkey.txt"
}

function configure_node() {
  NODE_HOME=$1
  NODE_MONIKER=$2

  echo "[${NODE_MONIKER}] Configuring app.toml and config.toml"

  APP_TOML="${NODE_HOME}/config/app.toml"
  CONFIG_TOML="${NODE_HOME}/config/config.toml"

  sed -i $SED_EXT 's/minimum-gas-prices = "0uplc"/minimum-gas-prices = "5uplc"/g' "${APP_TOML}"
  sed -i $SED_EXT 's/enable = false/enable = true/g' "${APP_TOML}"

  sed -i $SED_EXT 's|laddr = "tcp://127.0.0.1:26657"|laddr = "tcp://0.0.0.0:26657"|g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's|addr_book_strict = true|addr_book_strict = false|g' "${CONFIG_TOML}"

  sed -i $SED_EXT 's/timeout_propose = "3s"/timeout_propose = "500ms"/g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/timeout_prevote = "1s"/timeout_prevote = "500ms"/g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/timeout_precommit = "1s"/timeout_precommit = "500ms"/g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/timeout_commit = "5s"/timeout_commit = "500ms"/g' "${CONFIG_TOML}"
}

function configure_genesis() {
  NODE_HOME=$1
  NODE_MONIKER=$2

  echo "[${NODE_MONIKER}] Configuring genesis"

  GENESIS="${NODE_HOME}/config/genesis.json"
  GENESIS_TMP="${NODE_HOME}/config/genesis_tmp.json"

  # Default denom
  sed -i $SED_EXT 's/"stake"/"uplc"/' "${GENESIS}"

  # Short voting period
  sed -i $SED_EXT 's/"voting_period": "172800s"/"voting_period": "600s"/' "${GENESIS}"

  # Test accounts
  BASE_ACCOUNT_1="platine1rnr5jrt4exl0samwj0yegv99jeskl0hs4fya6d"
  # Mnemonic: sketch mountain erode window enact net enrich smoke claim kangaroo another visual write meat latin bacon pulp similar forum guilt father state erase bright
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_1}'", "coins": [{"denom": "uplc", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_1}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_2="platine1l9sq0se0jd3vklyrrtjchx4ua47awug5lzcxxq"
  # Mnemonic: ugly dirt sorry girl prepare argue door man that manual glow scout bomb pigeon matter library transfer flower clown cat miss pluck drama dizzy
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_2}'", "coins": [{"denom": "uplc", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_2}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_3="platine14wymr6h8a0rvynwayytssy5rhn3evltcj5tsaq"
  # Mnemonic: fix wheel picnic about army scan table fence device trust alter erupt wear donkey wood slender gold reunion grant quiz absurd tragic reform attitude
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_3}'", "coins": [{"denom": "uplc", "amount": "100"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_3}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_4="platine1c8tq9wcfwtunkunng5pe442jlmwns9sgltf4y7"
  # Mnemonic: horn slim pigeon winner capable piano soul find ignore crawl arrow genuine magnet nasty basic lamp scissors treat stick arm dress elbow trash naive
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_4}'", "coins": [{"denom": "uplc", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_4}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_5="platine1xnx6fe4fg4etlaxm0ty25j3ae5thhadrwz27qm"
  # Mnemonic: blue town hobby lens hawk deputy father tissue state choose another liquid license start push iron limb visa taste mother cause history tackle fiber
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_5}'", "coins": [{"denom": "uplc", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_5}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_6="platine1makysw7c788c7fzkfzn7al80wyfl4at6pewmhz"
  # Mnemonic: gallery hospital vicious demand orient piano melody vanish remind pistol elephant bracket olive kitten caution apart capital protect junior endorse run drama tiny patrol
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_6}'", "coins": [{"denom": "uplc", "amount": "100"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_6}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_VESTING_ACCOUNT="platine1lkqddnapqvz2hujx2trpj7xj6c9hmuq709rsy6"
  # Mnemonic: coach index fence broken very cricket someone casino dial truth fitness stay habit such three jump exotic spawn planet fragile walk enact angry great
  # shellcheck disable=SC2089
  BASE_VESTING_COIN="{\"denom\":\"uplc\",\"amount\":\"10001000000000000\"}"
  jq '.app_state.bank.balances += [{"address": "'${BASE_VESTING_ACCOUNT}'", "coins": [{"denom": "uplc", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.BaseVestingAccount", "base_account": {"address": "'${BASE_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1672531199"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  CONTINUOUS_VESTING_ACCOUNT="platine1353p46macvn444rupg2jstmx3tmz657ych5q2g"
  # Mnemonic: phone worry flame safe panther dirt picture pepper purchase tiny search theme issue genre orange merit stove spoil surface color garment mind chuckle image
  jq '.app_state.bank.balances += [{"address": "'${CONTINUOUS_VESTING_ACCOUNT}'", "coins": [{"denom": "uplc", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.ContinuousVestingAccount", "base_vesting_account": { "base_account": {"address": "'${CONTINUOUS_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1672531199"}, "start_time": "1630352459"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  DELAYED_VESTING_ACCOUNT="platine1njwu33lek5jt4kzlmljkp366ny4qpqusw9amu9"
  # Mnemonic: pilot text keen deal economy donkey use artist divide foster walk pink breeze proud dish brown icon shaft infant level labor lift will tomorrow
  jq '.app_state.bank.balances += [{"address": "'${DELAYED_VESTING_ACCOUNT}'", "coins": [{"denom": "uplc", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.DelayedVestingAccount", "base_vesting_account": { "base_account": {"address": "'${DELAYED_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1672531199"}}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  PERIODIC_VESTING_ACCOUNT="platine1uyngr0l3xtyj07js9sdew9mk50tqeq8lm9t8k5"
  # Mnemonic: want merge flame plate trouble moral submit wing whale sick meat lonely yellow lens enable oyster slight health vast weird radar mesh grab olive
  jq '.app_state.bank.balances += [{"address": "'${PERIODIC_VESTING_ACCOUNT}'", "coins": [{"denom": "uplc", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.PeriodicVestingAccount", "base_vesting_account": { "base_account": {"address": "'${PERIODIC_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1672531199"}, "start_time": "1672531179", "vesting_periods": [{"length": "20", "amount": ['"${BASE_VESTING_COIN}"']}]}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
}


NETWORK_CONFIG_DIR="network-config"
rm -rf $NETWORK_CONFIG_DIR
mkdir $NETWORK_CONFIG_DIR


# Generating node configurations
for ((i=1 ; i<=VALIDATORS_COUNT ; i++))
do
  NODE_MONIKER="validator-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  init_node "${NODE_HOME}" "${NODE_MONIKER}"
  configure_node "${NODE_HOME}" "${NODE_MONIKER}"
done

for ((i=1 ; i<=SEEDS_COUNT ; i++))
do
  NODE_MONIKER="seed-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  init_node "${NODE_HOME}" "${NODE_MONIKER}"
  configure_node "${NODE_HOME}" "${NODE_MONIKER}"
done

for ((i=1 ; i<=OBSERVERS_COUNT ; i++))
do
  NODE_MONIKER="observer-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  init_node "${NODE_HOME}" "${NODE_MONIKER}"
  configure_node "${NODE_HOME}" "${NODE_MONIKER}"
done


# Generating genesis
TMP_NODE_MONIKER="tmp"
TMP_NODE_HOME="${NETWORK_CONFIG_DIR}/${TMP_NODE_MONIKER}"
init_node "${TMP_NODE_HOME}" "${TMP_NODE_MONIKER}"
configure_genesis "${TMP_NODE_HOME}" "${TMP_NODE_MONIKER}"

mkdir "${TMP_NODE_HOME}/config/gentx"


# Adding genesis validators
for ((i=1 ; i<=VALIDATORS_COUNT ; i++))
do
  NODE_MONIKER="validator-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"

  platined keys add "operator-$i" --keyring-backend "test" --home "${NODE_HOME}"
  platined add-genesis-account "operator-$i" 20000000000000000uplc --keyring-backend "test" --home "${NODE_HOME}"

  NODE_ID=$(platined tendermint show-node-id --home "${NODE_HOME}")
  NODE_VAL_PUBKEY=$(platined tendermint show-validator --home "${NODE_HOME}")
  platined gentx "operator-$i" 1000000000000000uplc --chain-id "${CHAIN_ID}" --node-id "${NODE_ID}" \
    --pubkey "${NODE_VAL_PUBKEY}" --keyring-backend "test"  --home "${NODE_HOME}"

  cp "${NODE_HOME}/config/genesis.json" "${TMP_NODE_HOME}/config/genesis.json"
  cp -R "${NODE_HOME}/config/gentx/." "${TMP_NODE_HOME}/config/gentx"
done


echo "Collecting gentxs"
platined collect-gentxs --home "${TMP_NODE_HOME}"
platined validate-genesis --home "${TMP_NODE_HOME}"

# Distribute final genesis
for ((i=1 ; i<=VALIDATORS_COUNT ; i++))
do
  NODE_MONIKER="validator-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"
done

for ((i=1 ; i<=SEEDS_COUNT ; i++))
do
  NODE_MONIKER="seed-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"
done

for ((i=1 ; i<=OBSERVERS_COUNT ; i++))
do
  NODE_MONIKER="observer-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"
done

# Leave one copy of genesis in the root of network-config
cp "${TMP_NODE_HOME}/config/genesis.json" "${NETWORK_CONFIG_DIR}/genesis.json"

# Generate seeds.txt
SEEDS_STR=""

for ((i=1 ; i<=SEEDS_COUNT ; i++))
do
  NODE_MONIKER="seed-$i"
  NODE_P2P_PORT="26656"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  if ((i != 1))
  then
    SEEDS_STR="${SEEDS_STR},"
  fi

  SEEDS_STR="${SEEDS_STR}$(cat "${NODE_HOME}/node_id.txt")@${NODE_MONIKER}:${NODE_P2P_PORT}"
done

echo "${SEEDS_STR}" > "${NETWORK_CONFIG_DIR}/seeds.txt"

# We don't need the tmp node anymore
rm -rf "${TMP_NODE_HOME}"

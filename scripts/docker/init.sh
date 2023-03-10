#!/bin/sh

MONIKER=${MONIKER:-node001}
STAKE=${STAKE_TOKEN:-uplc}
CHAIN_ID=${CHAIN_ID:-platine-1}
CONFIG_DIR=${CONFIG_DIR:-$HOME/.platine}
BIN="${BINARY:-platined} --home $CONFIG_DIR"
PASSWORD=${PASSWORD:-platinepwd}

VALIDATOR_NMEMONIC=${VALIDATOR_NMEMONIC:-"midnight donkey dose rebel chunk bullet shed library aspect holiday control clump"}
DEV01_MNEMONIC=${DEV01_MNEMONIC:-"actual accuse plastic supply favorite banner trial company cloud wasp enable cactus"}
DEV02_MNEMONIC=${DEV02_MNEMONIC:-"virus rotate empty amused cherry shallow spell cruise lady lottery endorse tribe"}
DEV03_MNEMONIC=${DEV03_MNEMONIC:-"eagle escape crew record shield today minor choice funny void movie top"}
DEV04_MNEMONIC=${DEV04_MNEMONIC:-"issue invite only frog luxury square bus drink torch slow tomorrow head"}

# rm -rf $CONFIG_DIR
mkdir -p $CONFIG_DIR
if [ ! -f "$CONFIG_DIR/config/genesis.json" ] ; then
    $BIN init $MONIKER --chain-id $CHAIN_ID  >/dev/null 2>&1

    # modify genesis.json parameter
    sed -i "s/\"stake\"/\"$STAKE\"/" ${CONFIG_DIR}/config/genesis.json
    sed -i 's/"time_iota_ms": "1000"/"time_iota_ms": "1000"/' ${CONFIG_DIR}/config/genesis.json
    sed -i '/API server should be enabled/{n; s/enable = false/ enable = true/}; s/^swagger = .*/swagger = true/' $CONFIG_DIR/config/app.toml
    sed -i '/Rosetta API server should be enabled./{n; s/enable = true/ enable = false/};' $CONFIG_DIR/config/app.toml

    # add validator
    if ! $BIN keys show validator >/dev/null 2>&1; then
        # (echo "$PASSWORD"; echo "$PASSWORD") | $BIN keys add validator 
	echo VALIDATOR_NMEMONIC $VALIDATOR_NMEMONIC
        (echo $VALIDATOR_NMEMONIC; echo $PASSWORD; echo $PASSWORD ) | $BIN keys add validator --recover >/dev/null 2>&1;
	echo DEV01_MNEMONIC $DEV01_MNEMONIC
        (echo $DEV01_MNEMONIC; echo $PASSWORD; echo $PASSWORD ) | $BIN keys add dev01 --recover >/dev/null 2>&1;
	echo DEV02_MNEMONIC $DEV02_MNEMONIC
        (echo $DEV02_MNEMONIC; echo $PASSWORD; echo $PASSWORD ) | $BIN keys add dev02 --recover >/dev/null 2>&1;
	echo DEV03_MNEMONIC $DEV03_MNEMONIC
        (echo $DEV03_MNEMONIC; echo $PASSWORD; echo $PASSWORD ) | $BIN keys add dev03 --recover  >/dev/null 2>&1;
	echo DEV04_MNEMONIC $DEV04_MNEMONIC
        (echo $DEV04_MNEMONIC; echo $PASSWORD; echo $PASSWORD ) | $BIN keys add dev04 --recover  >/dev/null 2>&1;
    fi

    echo "$PASSWORD" | $BIN add-genesis-account $(echo "$PASSWORD"|$BIN keys show -a validator) "1000000000000000$STAKE" 
    echo "$PASSWORD" | $BIN add-genesis-account $(echo "$PASSWORD"|$BIN keys show -a dev01) "1000000000000000$STAKE" 
    echo "$PASSWORD" | $BIN add-genesis-account $(echo "$PASSWORD"|$BIN keys show -a dev02) "1000000000000000$STAKE" 
    echo "$PASSWORD" | $BIN add-genesis-account $(echo "$PASSWORD"|$BIN keys show -a dev03) "1000000000000000$STAKE" 
    echo "$PASSWORD" | $BIN add-genesis-account $(echo "$PASSWORD"|$BIN keys show -a dev04) "1000000000000000$STAKE" 

    #
    for addr in "$@"; do
      echo $addr
      $BIN add-genesis-account "$addr" "1000000000000000$STAKE"
    done

    (echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | $BIN gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID" --amount="250000000$STAKE" >/dev/null 2>&1

    $BIN collect-gentxs >/dev/null 2>&1
    $BIN validate-genesis --log_level warn
fi

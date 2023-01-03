#!/bin/sh
CONFIG_DIR=${CONFIG_DIR:-$HOME/.platine}
BIN="${BIN:-platined} --home $CONFIG_DIR"
TRACE=${1:-"--trace"}

CMD="$BIN start --rpc.laddr tcp://0.0.0.0:26657 $TRACE"

echo "Start: $CMD"

`$CMD`

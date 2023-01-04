#!/bin/sh

BASE_DIR=$(dirname "$0")

$BASE_DIR/init.sh "$@"
$BASE_DIR/start.sh

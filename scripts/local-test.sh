#!/bin/sh

BASE_DIR=$(dirname "$0")

$BASE_DIR/docker/init.sh "$@"
$BASE_DIR/docker/start.sh "--log_level warn"

#!/bin/sh

SCRIPT_DIR=$(dirname "$(realpath "$0")")

mkdir -p $SCRIPT_DIR/../bin

go build -o $SCRIPT_DIR/../bin/token-tracker

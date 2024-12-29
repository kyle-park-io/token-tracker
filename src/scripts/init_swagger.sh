#!/bin/sh

SCRIPT_DIR=$(dirname "$(realpath "$0")")

swag init --output ${SCRIPT_DIR}/../docs/blocktimestamp --parseDependency
swag init --output ${SCRIPT_DIR}/../docs/transfertracker --parseDependency

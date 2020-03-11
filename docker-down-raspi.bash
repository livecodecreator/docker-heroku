#!/usr/bin/env bash

set -eo pipefail
cd $(dirname $0)/cmd/raspi
export -p BASE_PATH=$(pwd)
export -p PATH=$BASE_PATH:$PATH

docker-compose down -v

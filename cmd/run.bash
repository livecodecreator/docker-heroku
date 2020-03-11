#!/usr/bin/env bash

set -eo pipefail
cd $(dirname $0)
export BASE_PATH=$(pwd)

go get -v github.com/gorilla/mux
go run -v github.com/livecodecreator/docker-heroku

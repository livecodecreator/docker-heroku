#!/usr/bin/env bash

set -eo pipefail
cd $(dirname $0)
export BASE_PATH=$(pwd)

heroku login
heroku deploy

#!/usr/bin/env bash

v=$(printf 'TEST%.0s' {1..5000})

while :; do
    curl -X POST -d "$v" http://localhost:8080/
done

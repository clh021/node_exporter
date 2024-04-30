#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
SCRIPT_PATH=$(readlink -f "$0")
cd "$(dirname "$(dirname "$SCRIPT_PATH")")" || exit 1
pwd
# CGO_ENABLED=0 go run ./node_exporter.go
promu --config="./hostinfo/.promu.yml" build --prefix ./

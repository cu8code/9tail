#!/bin/sh

set -e

if pgrep -f nats-server >/dev/null; then
    echo "NATS already running, skipping"
else
    nats-server -js &
fi

go run -mod=mod -tags=dev ./cmd/server/main.go
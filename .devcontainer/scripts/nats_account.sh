#!/bin/bash

set -e

DIR="$( dirname -- "${BASH_SOURCE[0]}"; )";
echo "DIR IS $DIR"
DEVCONTAINER_DIR="$DIR/.."

sudo chown -Rh $USER:$USER $DEVCONTAINER_DIR/nsc

echo "Dumping NATS user creds file"
sudo nsc --data-dir=$DEVCONTAINER_DIR/nsc/nats/nsc/stores generate creds -a LOC -n USER > /tmp/user.creds

echo "Dumping NATS sys creds file"
sudo nsc --data-dir=$DEVCONTAINER_DIR/nsc/nats/nsc/stores generate creds -a SYS -n sys > /tmp/sys.creds

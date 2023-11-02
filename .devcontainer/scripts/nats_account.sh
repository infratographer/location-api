#!/bin/bash

set -e
sudo chown -Rh vscode:vscode /workspaces/location-api/.devcontainer/nsc

echo "Dumping NATS user creds file"
nsc --data-dir=/workspaces/location-api/.devcontainer/nsc/nats/nsc/stores generate creds -a LOC -n USER > /tmp/user.creds

echo "Dumping NATS sys creds file"
nsc --data-dir=/workspaces/location-api/.devcontainer/nsc/nats/nsc/stores generate creds -a SYS -n sys > /tmp/sys.creds

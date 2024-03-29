#!/bin/sh
# script to bootstrap a nats operator environment

set -e

if nsc describe operator; then
    echo "operator exists, not overwriting config"
    exit 0
fi

echo "Cleaning up NATS environment"
rm -rf /nsc/*

echo "Creating NATS operator"
nsc add operator --generate-signing-key --sys --name LOCAL
nsc edit operator -u 'nats://nats:4222'
nsc list operators
nsc describe operator

export OPERATOR_SIGNING_KEY_ID=`nsc describe operator -J | jq -r '.nats.signing_keys | first'`

echo "Creating NATS account for location-api"
nsc add account -n LOC -K ${OPERATOR_SIGNING_KEY_ID}
nsc edit account LOC --sk generate --js-mem-storage -1 --js-disk-storage -1 --js-streams -1 --js-consumer -1
nsc describe account LOC

export ACCOUNTS_SIGNING_KEY_ID=`nsc describe account LOC -J | jq -r '.nats.signing_keys | first'`

echo "Creating NATS user for location-api"
nsc add user -n USER -K ${ACCOUNTS_SIGNING_KEY_ID}
nsc describe user USER

echo "Generating NATS resolver.conf"
nsc generate config --mem-resolver --sys-account SYS --config-file /nats/resolver.conf --force

#!/bin/bash
set -a
source ../infra/.mongo.env
set +a

sed -e "s/PLACEHOLDER_USER/${MONGO_USERNAME}/g" \
    -e "s/PLACEHOLDER_PASSWORD/${MONGO_PASSWORD}/g" \
    -e "s/PLACEHOLDER_DATABASE/${MONGO_DATABASE}/g" \
    -e "s/PLACEHOLDER_PORT/${MONGO_PORT}/g" \
    config.template.toml > config.toml
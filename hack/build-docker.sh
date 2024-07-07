#!/bin/bash

cd ../go-server

docker buildx build \
--push \
--platform linux/amd64,linux/arm64 \
--tag karlyan/go-server:latest .

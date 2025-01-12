#!/bin/bash

DOCKER_TAG=karlyan/sshd:2025.01

docker buildx build \
--push \
--platform linux/arm/v7,linux/arm64/v8,linux/amd64 \
--tag ${DOCKER_TAG} -f Dockerfile .

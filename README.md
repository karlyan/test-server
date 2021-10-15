# test-server

This project contains 2 test servers for testing http request
* go-server: go based server for http testing
* python-server: python based server for http testing

Both of them have Dockerfile for building docker images.
* go-server: size of 15.1MB docker image
* python-server: size of 56.6MB docker image
* go-server-tools: add some basic debug tools and awscli, helm, kubectl for easy debug in k8s. size: 534MB
* k8s-tools: some basic tools to support k8s

[![k8s-tools CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-k8s-tools.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-k8s-tools.yml)

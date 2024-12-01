# test-server

This project contains 2 test servers for testing http request
* go-server: go based server for http testing
* python-server: python based server for http testing

Both of them have Dockerfile for building docker images.
| Tools | status | Note |
| ------------- | ------------- | ------------- |
| k8s-tools | [![k8s-tools](https://github.com/karlyan/test-server/actions/workflows/docker-image-k8s-tools.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-k8s-tools.yml)  | some basic tools to support k8s size 380MB
| go-server | [![go-server CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server.yml) | size of 15.1MB docker image |
| python-server | [![python-server CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-python-server.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-python-server.yml) | size of 56.6MB docker image |
| go-server-tools | [![go-server-tools CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server-tools.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server-tools.yml) | add some basic debug tools and awscli, helm, kubectl for easy debug in k8s. size: 580MB |

This project also provides a generic helm chart for deploying the servers in k8s. People can use chart values driven deployment to deploy the servers in k8s.

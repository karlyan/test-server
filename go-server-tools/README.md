# go-server-tools

[`go-server`](../go-server/) plus a Kubernetes / AWS debugging toolchain layered on top:

- `kubectl`
- `helm`
- `eksctl`
- `aws-cli`
- `bash`, `curl`, `jq`, `yq`, `openssl`, `openssh`, `sshpass`, `groff`, `build-base`

Image size: ~580 MB. Multi-arch (`linux/amd64`, `linux/arm64`). The image still runs the `go-server` HTTP service on `:443` by default.

Use this when you want a single pod that both responds to test traffic *and* can shell out to inspect / mutate the cluster.

## Build dependency

```
karlyan/go-server  ──►  karlyan/go-server-tools
       (FROM)              (extends with tools)
```

When upgrading shared bases (Go, alpine, gin), bump and **republish `karlyan/go-server` first**, otherwise `go-server-tools` will rebuild on the old base. The `go-server CI` workflow has to be dispatched and the new tag pushed to Docker Hub before dispatching `go-server-tools CI`.

## Build

```bash
docker buildx build --push \
  --platform linux/amd64,linux/arm64 \
  --build-arg EKSCTL_VERSION=0.226 \
  --build-arg KUBECTL_VERSION=v1.36.0 \
  --build-arg HELM_VERSION=v4.1.4 \
  --tag karlyan/go-server-tools:latest \
  -f Dockerfile .
```

CI does the equivalent via the manually-triggered `go-server-tools CI` workflow.

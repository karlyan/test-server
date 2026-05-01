# k8s-tools

A small Alpine-based image with the tools you usually want in a Kubernetes debug pod:

- `kubectl` (current pin: `v1.36.0`)
- `helm` (current pin: `v4.1.4`)
- `aws-cli`
- `bash`, `curl`, `wget`, `jq`, `yq`, `openssl`, `gettext`, `coreutils`, `util-linux`

Image size: ~380 MB. Multi-arch (`linux/amd64`, `linux/arm64`).

## Run

```bash
docker run --rm -it karlyan/k8s-tools
```

In a cluster, drop it in as a debug pod:

```bash
kubectl run debug --rm -it --image=karlyan/k8s-tools --restart=Never -- bash
```

## Versions

`kubectl` and `helm` are installed under versioned filenames (`kubectl-v1.36.0`) with a stable symlink (`kubectl`). To pin a different version at build time:

```bash
docker buildx build --push \
  --platform linux/amd64,linux/arm64 \
  --build-arg KUBECTL_VERSION=v1.35.0 \
  --build-arg HELM_VERSION=v4.0.0 \
  --tag karlyan/k8s-tools:custom \
  -f Dockerfile .
```

CI does the equivalent via the manually-triggered `k8s-tools` workflow.

## Compared to `go-server-tools`

[`go-server-tools`](../go-server-tools/) bundles the same Kubernetes/AWS toolchain *into* the running `go-server` image, useful when you want one container that both serves traffic and can poke around the cluster. `k8s-tools` is the standalone shell-only variant.

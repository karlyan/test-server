# server-tools

A general-purpose Ubuntu LTS image preloaded with networking and SSH tooling — useful as an interactive box on a host or in a cluster.

Base: `ubuntu:resolute` (26.04 LTS).

Includes: `curl`, `git`, `gnupg`, `iproute2`, `iputils-ping`, `net-tools`, `openssh-server`, `sudo`, `vim`, `wget`.

## Build

There is no GitHub Actions workflow for this image — it is built and tagged locally:

```bash
./build.sh
# equivalent to:
# docker build -f Dockerfile -t karlyan/server-tools:latest .
```

## Compared to other images

| Use | Image |
| --- | --- |
| Kubernetes ops (kubectl, helm, aws-cli) | [`k8s-tools`](../k8s-tools/) |
| go-server + Kubernetes/AWS tools | [`go-server-tools`](../go-server-tools/) |
| Just SSH access to a host | [`sshd`](../sshd/) |
| General Linux toolbox (this image) | `server-tools` |

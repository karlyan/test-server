# test-server

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![go-server CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server.yml)
[![python-server CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-python-server.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-python-server.yml)
[![k8s-tools CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-k8s-tools.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-k8s-tools.yml)
[![go-server-tools CI](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server-tools.yml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/docker-image-go-server-tools.yml)
[![Lint and Test Charts](https://github.com/karlyan/test-server/actions/workflows/lint-test.yaml/badge.svg)](https://github.com/karlyan/test-server/actions/workflows/lint-test.yaml)

A small monorepo of HTTP test servers, debugging containers, and a generic Helm chart used to deploy them on Kubernetes.

## Components

| Path | Image | Description | Size |
| --- | --- | --- | --- |
| [`go-server/`](go-server/) | `karlyan/go-server` | Minimal Gin HTTP server (`/` echoes client IP, `/debug` dumps the full request) | ~15 MB |
| [`python-server/`](python-server/) | `karlyan/python-server` | Minimal Flask HTTP server | ~57 MB |
| [`k8s-tools/`](k8s-tools/) | `karlyan/k8s-tools` | Alpine + `kubectl`, `helm`, `aws-cli`, `jq`, `yq`, `curl` | ~380 MB |
| [`go-server-tools/`](go-server-tools/) | `karlyan/go-server-tools` | `go-server` + Kubernetes/AWS debug tools (`kubectl`, `helm`, `eksctl`, `aws-cli`) | ~580 MB |
| [`server-tools/`](server-tools/) | `karlyan/server-tools` | Ubuntu LTS + general networking & SSH tooling | – |
| [`sshd/`](sshd/) | `karlyan/sshd` | Alpine + OpenSSH on port 2233 (key-based) | – |
| [`charts/generic-chart/`](charts/generic-chart/) | – | Values-driven generic Helm chart (Deployment, Service, Ingress, Gateway API, HPA, RBAC, cert-manager, NetworkPolicy, storage) | – |

## Quick start

Run any server locally with Docker:

```bash
docker run --rm -p 8443:443 karlyan/go-server          # http://localhost:8443
docker run --rm -p 4999:4999 karlyan/python-server     # http://localhost:4999
docker run --rm -it karlyan/k8s-tools                  # interactive shell
```

Deploy `go-server` on Kubernetes via the generic chart:

```bash
helm upgrade --install --create-namespace -n test-server test-server generic-chart \
  --repo "https://test-server.github.yyzd.me" --version "^1.0.0" -f - <<EOF
deployment:
  enabled: true
  image:
    repository: karlyan/go-server
    tag: latest
service:
  enabled: true
  ports:
    http:
      enabled: true
      protocol: TCP
      containerPort: 443
      servicePort: 443
EOF
```

See [`charts/generic-chart/README.md`](charts/generic-chart/README.md) for the full values reference.

## Repository layout

The idea behind the generic chart is to avoid one chart per service. All workload-specific differences are injected through values from a CD system, with per-environment value files kept in separate Git repos.

```
.
├── charts/generic-chart/   # Single Helm chart used by every service
├── go-server/              # Source + Dockerfile for the Gin server
├── python-server/          # Source + Dockerfile for the Flask server
├── go-server-tools/        # Debug image, FROM karlyan/go-server
├── k8s-tools/              # Standalone Kubernetes ops image
├── server-tools/           # Standalone networking/debug image
├── sshd/                   # Standalone OpenSSH image
└── .github/workflows/      # CI for image builds and chart release
```

## Development

| Task | Command |
| --- | --- |
| Build go-server locally | `cd go-server && go build .` |
| Run go-server locally | `cd go-server && go run .` (listens on `:443`) |
| Build python-server locally | `cd python-server && uv sync` |
| Run python-server locally | `cd python-server && uv run -m app` (listens on `:4999`) |
| Lint the Helm chart | `ct lint --config .github/linters/ct.yaml` |
| Run chart unit tests | `helm unittest --strict charts/generic-chart` |
| Render the chart | `helm template charts/generic-chart -f some-values.yaml` |

The chart unit tests use [helm-unittest](https://github.com/helm-unittest/helm-unittest):

```bash
helm plugin install --verify=false https://github.com/helm-unittest/helm-unittest --version v1.0.3
```

## CI / release flow

* **Image workflows** (`docker-image-*.yml`) are **manual** (`workflow_dispatch`). Pushing code does not rebuild images. Trigger via the Actions tab or `gh workflow run "go-server CI" -f tagName=latest`. Multi-arch (`amd64`, `arm64`) push to Docker Hub `karlyan/<name>:<tagName>`.
* **`lint-test.yaml`** runs `ct lint` and `helm unittest` on every PR and on `main` pushes that touch the chart.
* **`release-chart.yaml`** packages the chart and publishes a GitHub Release whenever `charts/**` changes on `main`. `chart-releaser-action` uses `skip_existing: true`, so **`Chart.yaml`'s `version` must be bumped** for every chart change or the release is silently skipped.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[Apache 2.0](LICENSE).

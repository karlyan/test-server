# generic-chart

A single, values-driven Helm chart used to deploy every service in this repo (and beyond). Instead of writing a chart per service, all workload-specific differences are injected via values from the CD system.

## Resources supported

Each top-level section is gated by an `enabled` flag (default `false`) — turn on what you need:

| Section | Resource(s) |
| --- | --- |
| `deployment` | `Deployment` |
| `service` | `Service` (named-port map) |
| `serviceGeneric` | Second free-form `Service` (raw `spec`) |
| `ingress` | `Ingress` |
| `gateways` | `Gateway` (Gateway API) — **map by key**, one per entry |
| `httpRoutes` | `HTTPRoute` (Gateway API) — **map by key**, one per entry |
| `autoscaling` | `HorizontalPodAutoscaler` |
| `serviceAccount` / `role` / `roleBinding` / `clusterRole` / `clusterRoleBinding` | RBAC |
| `secret` | `Secret` |
| `certificate` / `clusterIssuer` | `cert-manager.io/v1` |
| `pvc` / `pv` / `storageClass` | Storage |
| `networkPolicy` | `NetworkPolicy` |

`gateways` and `httpRoutes` are **maps** (not single instances) because typical multi-host setups need one HTTPRoute per hostname → backend pair. Each enabled key renders one resource named `{fullname}-{key}` (override via `.name`).

## Quick start

Install from the published chart repo:

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

A more complete example (Deployment + Service + multi-host HTTPRoute) is in the comments at the top of [`values.yaml`](values.yaml).

## Local development

```bash
# Lint (mirrors what CI runs)
ct lint --config ../../.github/linters/ct.yaml

# Render with sample values
helm template charts/generic-chart -f some-values.yaml

# Run unit tests
helm plugin install https://github.com/helm-unittest/helm-unittest --version v1.0.3   # once
helm unittest --strict .
```

Tests live under [`tests/`](tests/), one suite per template (`*_test.yaml`). New templates **must** ship with a matching suite — CI fails otherwise.

## Releasing

`chart-releaser-action` runs on `main` whenever `charts/**` changes. It uses `skip_existing: true`, so:

> **Bump `Chart.yaml`'s `version` for every chart change**, otherwise the new package is silently dropped.

CI's `ct lint` enforces this on PRs (`version-bump` check fails if you forget).

The published chart repo is at `https://test-server.github.yyzd.me` (a personal GitHub Pages mirror).

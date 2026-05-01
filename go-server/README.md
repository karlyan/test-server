# go-server

A minimal HTTP test server written in Go using [Gin](https://github.com/gin-gonic/gin). Listens on `0.0.0.0:443`.

## Endpoints

| Path | Description |
| --- | --- |
| `GET /` | Plain-text response showing `Request.RemoteAddr` and the resolved client IP. |
| `GET /debug` | Plain-text dump of the entire incoming HTTP request (headers + body). |

Useful for inspecting how proxies / ingresses / load balancers rewrite traffic.

## Run

Locally:

```bash
go run .
```

Via Docker:

```bash
docker run --rm -p 8443:443 karlyan/go-server
curl -s http://localhost:8443/
```

## Build

The Dockerfile is a multi-stage `golang:1.26-alpine` → `alpine:3.23` build. Multi-arch image (`linux/amd64`, `linux/arm64`):

```bash
docker buildx build --push \
  --platform linux/amd64,linux/arm64 \
  --tag karlyan/go-server:latest \
  -f Dockerfile .
```

CI does the same via the manually-triggered `go-server CI` workflow.

## Dependents

[`go-server-tools`](../go-server-tools/) extends this image with Kubernetes / AWS debugging tools — bump and republish `karlyan/go-server` first when both change.

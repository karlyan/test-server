# python-server

A minimal HTTP test server written in Python using [Flask](https://flask.palletsprojects.com/). Listens on `0.0.0.0:4999`.

Counterpart to [`go-server`](../go-server/) — same purpose, different language, used to compare runtime / image-size characteristics.

## Endpoints

| Path | Description |
| --- | --- |
| `GET /` | Returns `Hello world`. |

## Run

Dependencies are managed with [uv](https://github.com/astral-sh/uv).

Locally:

```bash
uv sync
uv run python -m app
# or:
./run.sh
```

Via Docker:

```bash
docker run --rm -p 4999:4999 karlyan/python-server
curl -s http://localhost:4999/
```

## Build

The Dockerfile uses `python:3.14-slim-bookworm`, installs `uv` from the official installer, and runs the app under `supervisord` so the container correctly handles signals and restarts. Multi-arch image (`linux/amd64`, `linux/arm64`):

```bash
docker buildx build --push \
  --platform linux/amd64,linux/arm64 \
  --tag karlyan/python-server:latest \
  -f Dockerfile .
```

CI does the same via the manually-triggered `python-server CI` workflow.

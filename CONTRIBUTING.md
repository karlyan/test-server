# Contributing

Thanks for your interest. This is a small project, so there is no heavy process — but a few rules are worth following.

## Local checks before opening a PR

Whichever component you touched:

| Touching | Run |
| --- | --- |
| `go-server/` | `cd go-server && go build .` |
| `python-server/` | `cd python-server && uv sync && uv run python -m app` |
| `charts/generic-chart/` | `ct lint --config .github/linters/ct.yaml && helm unittest --strict charts/generic-chart` |
| Any image | Verify the Dockerfile builds with `docker buildx build --platform linux/amd64,linux/arm64 -f <path>/Dockerfile <path>` |

For the chart unit tests you need the helm-unittest plugin once:

```bash
helm plugin install --verify=false https://github.com/helm-unittest/helm-unittest --version v1.0.3
```

## Chart changes need a version bump

`chart-releaser-action` runs with `skip_existing: true`, so any change under `charts/**` **must** bump `charts/generic-chart/Chart.yaml`'s `version`, otherwise the new chart is silently dropped on the floor. CI's `ct lint` enforces this on PRs (the `version-bump` check fails if you forget).

When you add a new template, also add a matching suite under `charts/generic-chart/tests/<resource>_test.yaml` covering at minimum: disabled default renders nothing, enabled renders the right `kind` with the expected default name, and any non-trivial conditional logic.

## Image workflows are manual

`docker-image-*.yml` are `workflow_dispatch` only — pushing to `main` does **not** rebuild any image. After your PR merges, trigger the relevant workflow yourself:

```bash
gh workflow run "go-server CI" -f tagName=latest
```

Note the dependency: `go-server-tools` `FROM karlyan/go-server`. If both change, dispatch `go-server CI` first, wait for the new tag to appear on Docker Hub, then dispatch `go-server-tools CI`.

## Commit style

Look at `git log` and follow the existing tone — short imperative subject, optional body explaining *why*. Multi-component changes can land in a single commit if they belong to one logical change.

## License

By contributing, you agree your contribution is licensed under [Apache 2.0](LICENSE).

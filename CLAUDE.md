# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 仓库定位

这是一个用来做 HTTP 测试 / Kubernetes 调试的多镜像 monorepo。每个顶层目录是一个独立可发布的 Docker 镜像或 Helm chart：

- `go-server/` — Gin 写的最小 HTTP 服务器，监听 `0.0.0.0:443`，提供 `/`（回显客户端 IP）和 `/debug`（dump 整个请求）。镜像约 15MB。
- `python-server/` — Flask 应用，监听 `0.0.0.0:4999`，用 [uv](https://github.com/astral-sh/uv) 管理依赖，容器内由 `supervisord` 启动 `run.sh`（即 `uv run -m app`）。
- `go-server-tools/` — 在 `karlyan/go-server` 基础上叠加 awscli / eksctl / kubectl / helm，方便 K8s 内部调试。
- `k8s-tools/` — alpine + kubectl + helm + awscli + 常用工具（`bash curl jq yq` 等），约 380MB。
- `server-tools/` — Ubuntu jammy + 常用排查工具（`curl git net-tools openssh-server` 等）。
- `sshd/` — alpine + openssh，监听 2233 端口；`entry.sh` 处理 host key 生成和 authorized_keys 权限。
- `charts/generic-chart/` — 通用 Helm chart（详见下文）。
- `hack/build-docker.sh` — 本地手动多架构 build & push 的样例脚本。

## 构建 / 运行常用命令

### go-server
```bash
# 本地直接跑
cd go-server && go run .          # 监听 :443

# 多架构构建并推送（参考 hack/build-docker.sh）
docker buildx build --push --platform linux/amd64,linux/arm64 \
  --tag karlyan/go-server:latest go-server/
```

### python-server
```bash
cd python-server
uv sync                            # 安装依赖
uv run -m app                      # 本地启动，监听 :4999
# 或
./run.sh
```

### 镜像（任意目录下）
所有 Dockerfile 都依赖 `BUILDPLATFORM` / `TARGETOS` / `TARGETARCH`，必须用 `docker buildx`。各目录都是 build context 根。

### Helm chart
```bash
# 本地 lint（CI 跑的也是这个）
ct lint --config .github/linters/ct.yaml

# 单元测试（helm-unittest 插件，需 helm plugin install 一次）
helm unittest --strict charts/generic-chart

# 模板渲染检查
helm template charts/generic-chart -f some-values.yaml
```

测试文件在 `charts/generic-chart/tests/<resource>_test.yaml`，每个模板一个 suite。新增模板必须配套加 suite——CI 会跑 `helm unittest`。

## CI / 发布流程

- `.github/workflows/docker-image-*.yml` 都是 **`workflow_dispatch` 手动触发**，输入 `tagName` 决定镜像 tag（默认 `latest`），多架构（amd64+arm64）推送到 Docker Hub `karlyan/<name>`。改 Dockerfile 不会自动出新镜像，需要手动跑 workflow。
- `.github/workflows/release-chart.yaml` — 当 `charts/**` 在 `main` 分支有变更时，用 `helm/chart-releaser-action` 自动打包并发布到 GitHub Pages 的 chart repo（指向就是 README 里的 `https://github.yyzd.me`）。**修改 chart 时记得同步 bump `charts/generic-chart/Chart.yaml` 里的 `version`，否则 chart-releaser 会因为 `skip_existing: true` 跳过。**
- `.github/workflows/lint-test.yaml` — PR 触发 `ct lint` + `helm unittest`（强制，所有 chart 模板都有对应 suite）。

## 架构要点

### generic-chart 的设计意图

`charts/generic-chart` 是这个仓库的核心架构决策：**不为每个服务写一个 chart，而是用一个变量驱动的通用 chart**，所有差异通过 values 注入（在 CD 系统 / 各环境的 git repo 里维护）。

- 顶层每个能力（`deployment`, `service`, `ingress`, `gateways` / `httpRoutes`, `hpa`, `cert-manager`, `networkpolicy`, `rbac`, `secret`, `storage`）都有 `enabled` 开关，默认全 `false`，按需启用。
- `gateways` / `httpRoutes` 是 **map by key** 而非单实例（区别于 `ingress`），每个 key 渲染一个资源、名字默认 `{fullname}-{key}`——典型场景一个 app 多个 host 走不同 backend，必须靠多 HTTPRoute 区分。
- `deployment.ports` 是一个命名 map（`http`, `https`...），`livenessProbe` / `readinessProbe` 通过 `port: http` 这样的名字引用——新增端口时要保持名字一致，否则探针会指空。
- README 里的 helm 命令展示了典型用法：`--reuse-values` + 通过 stdin 传入服务特定的 values。

新增功能时优先扩展 `generic-chart`，不要新建 chart。

### 镜像之间的依赖

- `go-server-tools` 的 `FROM karlyan/go-server` 意味着先要发 `go-server` 才能 build `go-server-tools`。
- `k8s-tools` / `go-server-tools` 都把 `kubectl` / `helm` 安装为 `<bin>-<version>` 并通过 symlink 暴露——升级版本时改 `ENV *_VERSION` 即可，不要破坏 symlink 模式。

## 项目特定约定

- 服务监听端口：go-server `:443`，python-server `:4999`，sshd `:2233`。chart values 默认值与之对应。
- README 里的 chart repo 地址 `https://github.yyzd.me` 是用户自己的 GitHub Pages 镜像，不是公共仓库。
- 这是 **个人项目**（路径 `karlyan/`），git push 时使用 `karlyan` 账号；用 `gh auth status` 确认后再 push。

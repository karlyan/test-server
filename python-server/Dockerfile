FROM --platform=${BUILDPLATFORM} python:3.13-slim-bookworm

ENV DEBIAN_FRONTEND=noninteractive \
    TZ=Etc/UTC \
    PIP_NO_CACHE_DIR=1

ARG TARGETOS
ARG TARGETARCH

LABEL org.opencontainers.image.platform="${TARGETOS}/${TARGETARCH}"

RUN set -eux; \
    apt-get update; \
    apt-get upgrade -y --no-install-recommends; \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        tzdata \
        curl \
        supervisor; \
    \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime; echo $TZ > /etc/timezone; \
    export UV_INSTALL_DIR=/usr/local/bin; \
    curl -LsSf https://astral.sh/uv/install.sh | sh; \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

SHELL ["/bin/bash", "-c"]

COPY pyproject.toml poetry.lock* uv.lock* requirements*.txt* ./
RUN set -eux; \
    if [ -f "uv.lock" ] || [ -f "poetry.lock" ] || ls requirements*.txt >/dev/null 2>&1; then \
        uv sync --frozen || uv sync; \
    else \
        true; \
    fi

COPY . .
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 4999
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]

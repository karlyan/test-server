FROM alpine:3.14

# base packages
RUN apk add --no-cache --upgrade coreutils util-linux gettext bash curl wget jq yq openssl aws-cli

ARG KUBECTL_VERSION
ARG HELM_VERSION
ARG HELMFILE_VERSION

# kubectl
RUN echo "KUBECTL_VERSION=${KUBECTL_VERSION}" && \
    { [ -n "${KUBECTL_VERSION}" ] || KUBECTL_VERSION="v1.22.2"; } && \
    curl -LO "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    kubectl version --client

# helm 3
RUN echo "HELM_VERSION=${HELM_VERSION}" && \
    HELM_VERSION="${HELM_VERSION:-v3.7.1}" && \
    curl -sO https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    tar xf ./helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    mv -f ./linux-amd64/helm /usr/local/bin/helm && \
    helm version && \
    rm -rf ./helm-${HELM_VERSION}-linux-amd64.tar.gz ./linux-amd64

# helm 3.6.3
RUN echo "HELM_VERSION=${HELM_VERSION}" && \
    HELM_VERSION="${HELM_VERSION:-v3.6.3}" && \
    curl -sO https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    tar xf ./helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    mv -f ./linux-amd64/helm /usr/local/bin/helm-${HELM_VERSION} && \
    helm-${HELM_VERSION} version && \
    rm -rf ./helm-${HELM_VERSION}-linux-amd64.tar.gz ./linux-amd64

# helmfile
RUN echo "HELMFILE_VERSION=${HELMFILE_VERSION}" && \
    HELMFILE_VERSION="${HELMFILE_VERSION:-v0.141.0}" && \
    curl -sLO https://github.com/roboll/helmfile/releases/download/${HELMFILE_VERSION}/helmfile_linux_amd64 && \
    mv -f helmfile_linux_amd64 /usr/local/bin/helmfile && \
    chmod u+x /usr/local/bin/helmfile && \
    helmfile version

CMD ["/bin/bash"]
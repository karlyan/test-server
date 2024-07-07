FROM karlyan/go-server

# TARGETOS and TARGETARCH are set by the buildx build command
ARG TARGETOS TARGETARCH

# base packages
RUN apk add --no-cache --upgrade \
        bash coreutils build-base gettext bash curl jq yq openssl openssh sshpass groff aws-cli

# eksctl https://github.com/eksctl-io/eksctl/releases
ENV EKSCTL_VERSION=${EKSCTL_VERSION:-"0.176"}
RUN CMD_NAME=eksctl && \
    curl --silent --location "https://github.com/eksctl-io/eksctl/releases/download/v${EKSCTL_VERSION}.0/${CMD_NAME}_$(uname -s)_${TARGETARCH}.tar.gz" | tar xz -C /tmp && \
    chmod +x /tmp/${CMD_NAME} && mv /tmp/${CMD_NAME} /usr/local/bin/${CMD_NAME}-${EKSCTL_VERSION} && \
    ln -sf /usr/local/bin/${CMD_NAME}-${EKSCTL_VERSION} /usr/local/bin/${CMD_NAME} && \
    ${CMD_NAME} version

# kubectl
ENV KUBECTL_VERSION=${KUBECTL_VERSION:-"v1.30.0"}
RUN CMD_NAME=kubectl && \
    curl -sLO https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/${TARGETOS}/${TARGETARCH}/${CMD_NAME} && \
    chmod +x ./${CMD_NAME} && mv -f ./${CMD_NAME} /usr/local/bin/${CMD_NAME}-${KUBECTL_VERSION}

# helm 3
ENV HELM_VERSION=${HELM_VERSION:-"v3.14.4"}
RUN CMD_NAME=helm && \
    curl --silent --location "https://get.helm.sh/${CMD_NAME}-${HELM_VERSION}-${TARGETOS}-${TARGETARCH}.tar.gz" | tar xz -C /tmp && \
    chmod +x /tmp/${TARGETOS}-${TARGETARCH}/${CMD_NAME} && mv -f /tmp/${TARGETOS}-${TARGETARCH}/${CMD_NAME} /usr/local/bin/${CMD_NAME}-${HELM_VERSION}

EXPOSE 443
CMD ["/usr/local/bin/app"]
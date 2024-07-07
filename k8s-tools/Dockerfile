FROM --platform=$BUILDPLATFORM alpine:3.20

# TARGETOS and TARGETARCH are set by the buildx build command
ARG TARGETOS TARGETARCH

# base packages
RUN apk add --no-cache --upgrade coreutils util-linux gettext bash curl wget jq yq openssl aws-cli

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

CMD ["/bin/bash"]

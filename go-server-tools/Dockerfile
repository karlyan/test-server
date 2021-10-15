FROM karlyan/go-server

ENV EKSCTL_VERSION=v0.69.0
ENV KUBECTL_VERSION=v1.22.2
ENV HELM_VERSION=v3.7.1

# base packages
RUN apk add --no-cache --upgrade \
        bash coreutils build-base gettext bash curl jq yq openssl openssh sshpass groff

# aws-cli v2
ENV AWSCLI_VERSION=2.2.41
ENV GLIBC_VER=2.31-r0
RUN apk add --update --no-cache groff
RUN apk add --no-cache --virtual .dependencies binutils curl \
    && curl -sL https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub -o /etc/apk/keys/sgerrand.rsa.pub \
    && curl -sLO https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VER}/glibc-${GLIBC_VER}.apk \
    && curl -sLO https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VER}/glibc-bin-${GLIBC_VER}.apk \
    && curl -sLO https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VER}/glibc-i18n-${GLIBC_VER}.apk \
    && apk add --no-cache --virtual .glibc \
        glibc-${GLIBC_VER}.apk \
        glibc-bin-${GLIBC_VER}.apk \
        glibc-i18n-${GLIBC_VER}.apk \
    && /usr/glibc-compat/bin/localedef -i en_US -f UTF-8 en_US.UTF-8 \
    && curl -sL https://awscli.amazonaws.com/awscli-exe-linux-x86_64-${AWSCLI_VERSION}.zip -o awscliv2.zip \
    && unzip awscliv2.zip \
    && aws/install \
    && rm -rf \
        awscliv2.zip \
        aws \
        /usr/local/aws-cli/v2/*/dist/aws_completer \
        /usr/local/aws-cli/v2/*/dist/awscli/data/ac.index \
        /usr/local/aws-cli/v2/*/dist/awscli/examples \
        glibc-*.apk \
    && apk del --purge .dependencies

# eksctl
RUN curl -sSLO https://github.com/weaveworks/eksctl/releases/download/${EKSCTL_VERSION}/eksctl_Linux_amd64.tar.gz && \
    tar xf ./eksctl_Linux_amd64.tar.gz && \
    mv -f ./eksctl /usr/local/bin/eksctl && \
    eksctl version && \
    rm -rf ./eksctl_Linux_amd64.tar.gz

# kubectl
RUN curl -sSLO https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \
    mv -f ./kubectl /usr/local/bin/kubectl && \
    kubectl version --client

# helm 3
RUN curl -sSLO https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    tar xf ./helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    mv -f ./linux-amd64/helm /usr/local/bin/helm && \
    helm version && \
    rm -rf ./helm-${HELM_VERSION}-linux-amd64.tar.gz ./linux-amd64

EXPOSE 443
CMD ["/usr/local/bin/app"]
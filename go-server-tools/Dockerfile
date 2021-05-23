FROM python:3-alpine3.13 AS installer

ENV AWSCLI_VERSION=2.2.4

RUN apk add --no-cache \
    gcc \
    git \
    libc-dev \
    libffi-dev \
    openssl-dev \
    py3-pip \
    zlib-dev \
    make \
    cmake

RUN git clone --recursive  --depth 1 --branch ${AWSCLI_VERSION} --single-branch https://github.com/aws/aws-cli.git

WORKDIR /aws-cli

# Follow https://github.com/six8/pyinstaller-alpine to install pyinstaller on alpine
RUN pip install --no-cache-dir --upgrade pip \
    && pip install --no-cache-dir pycrypto \
    && git clone --depth 1 --single-branch --branch v$(grep PyInstaller requirements-build.txt | cut -d'=' -f3) https://github.com/pyinstaller/pyinstaller.git /tmp/pyinstaller \
    && cd /tmp/pyinstaller/bootloader \
    && CFLAGS="-Wno-stringop-overflow -Wno-stringop-truncation" python ./waf configure --no-lsb all \
    && pip install .. \
    && rm -Rf /tmp/pyinstaller \
    && cd - \
    && boto_ver=$(grep botocore setup.cfg | cut -d'=' -f3) \
    && git clone --single-branch --branch v2 https://github.com/boto/botocore /tmp/botocore \
    && cd /tmp/botocore \
    && git checkout $(git log --grep $boto_ver --pretty=format:"%h") \
    && pip install . \
    && rm -Rf /tmp/botocore  \
    && cd -

RUN sed -i '/botocore/d' requirements.txt \
    && scripts/installers/make-exe

RUN unzip dist/awscli-exe.zip && \
    ./aws/install --bin-dir /aws-cli-bin

FROM karlyan/go-server

ENV KUBECTL_VERSION=v1.21.1
ENV HELM_VERSION=v3.5.4

ENV PATH "$PATH:/root/.local/bin"

# base packages
RUN apk add --no-cache --upgrade \
        coreutils \
        build-base \
        gettext \
        bash \
        curl \
        jq \
        python3 \
        py3-pip \
        openssl \
        openssh sshpass

# aws-cli v2
COPY --from=installer /usr/local/aws-cli/ /usr/local/aws-cli/
COPY --from=installer /aws-cli-bin/ /usr/local/bin/

# eksctl
RUN curl -sSLO https://github.com/weaveworks/eksctl/releases/download/latest_release/eksctl_$(uname -s)_amd64.tar.gz && \
    tar xf ./eksctl_$(uname -s)_amd64.tar.gz && \
    mv -f ./eksctl /usr/local/bin/eksctl && \
    eksctl version && \
    rm -rf ./eksctl_$(uname -s)_amd64.tar.gz

# kubectl
RUN KUBECTL_VERSION="${KUBECTL_VERSION}" && \
    curl -sSLO https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \
    mv -f ./kubectl /usr/local/bin/kubectl && \
    kubectl version --client

# helm 3
RUN HELM_VERSION="${HELM_VERSION}" && \
    curl -sSLO https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    tar xf ./helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    mv -f ./linux-amd64/helm /usr/local/bin/helm && \
    helm version && \
    rm -rf ./helm-${HELM_VERSION}-linux-amd64.tar.gz ./linux-amd64

#yq utility
RUN curl -sSLo /usr/local/bin/yq "https://github.com/mikefarah/yq/releases/download/3.3.0/yq_linux_amd64" && \
    chmod +x ./usr/local/bin/yq

EXPOSE 443
CMD ["/usr/local/bin/app"]
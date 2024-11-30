# create a ubuntu container with jammy release
FROM ubuntu:jammy

# install basic tools and remove cache
RUN apt-get update && apt-get install -y \
    curl \
    git \
    gnupg \
    iproute2 \
    iputils-ping \
    net-tools \
    openssh-server \
    sudo \
    vim \
    wget \
    && rm -rf /var/lib/apt/lists/*




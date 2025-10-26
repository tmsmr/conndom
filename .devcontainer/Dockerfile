FROM debian:trixie

RUN apt-get update && apt-get install -y \
    sudo \
    build-essential \
    curl \
    git \
    wget \
    nano \
    golang \
    openssh-client \
    dbus-bin \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -ms /bin/bash developer \
    && usermod -aG sudo developer \
    && echo "developer ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

ENV GOPATH="/home/developer/.go" \
    PATH="$GOPATH/bin:/usr/local/go/bin:$PATH"

USER developer
WORKDIR /home/developer

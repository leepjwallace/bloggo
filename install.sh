#!/bin/bash

ARCH="$(uname -m)"
OS="$(uname | awk '{print tolower($0)}')"

if [ "$ARCH" = "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" = "aarch64" ]; then
    ARCH="arm64"
elif [ "$ARCH" = "armv7l" ]; then
    ARCH="arm"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

if [ "$OS" = "darwin" ]; then
    OS="mac"
elif [ "$OS" = "linux" ]; then
    OS="linux"
else
    echo "Unsupported OS: $OS"
    exit 1
fi


base_url="https://github.com/"
username="leepjwallace"
repo="/bloggo/releases/download/"
tag="latest"
file="/bloggo_${OS}_${ARCH}"

url=$base_url$username$repo$tag$file

curl -L $url -o bloggo

chmod 777 bloggo
sudo mv bloggo /usr/local/bin

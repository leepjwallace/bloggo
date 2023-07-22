#!/bin/bash

# Detect the OS (works for Linux and macOS)
OS=$(uname -s)

# Detect the architecture
ARCH=$(uname -m)

# Map the architecture string reported to a value used in your binary's name
if [ "${ARCH}" == "x86_64" ]; then
    ARCH="amd64"
elif [ "${ARCH}" == "i386" ]; then
    ARCH="386"
elif [ "${ARCH}" == "arm" ]; then
    ARCH="arm"  # or "arm64" if you're targeting that
else
    echo "Unsupported architecture: ${ARCH}"
    exit 1
fi

# Form the download URL
URL="https://github.com/leepjwallace/bloggo/releases/latest/bloggo_${OS,,}_${ARCH}"

# Download the binary
curl -Lo bloggo ${URL}

# Make it executable
chmod +x bloggo

# Move the binary to a location in the PATH
sudo mv bloggo /usr/local/bin

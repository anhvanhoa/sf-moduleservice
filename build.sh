#!/bin/bash

# Build script for Docker with private repository support
# Usage: ./build.sh [github_token]

GITHUB_TOKEN=${1:-$GITHUB_TOKEN}

if [ -z "$GITHUB_TOKEN" ]; then
    echo "Error: GitHub token is required"
    echo "Usage: ./build.sh <github_token>"
    echo "Or set GITHUB_TOKEN environment variable"
    exit 1
fi

echo "Building Docker image with GitHub token..."

docker build \
    --build-arg GITHUB_TOKEN="$GITHUB_TOKEN" \
    -t permission-service:latest \
    -f DockerFile \
    .

echo "Build completed!"

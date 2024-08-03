#!/bin/bash

TOKEN=$1
REPO_NAME=$2
TAG_NAME=$3
OS=$4
ARCH=$5

curl -L \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${TOKEN}" \
    -H "Content-Type: application/octet-stream" \
    --data-binary @./build/images-${OS}-${ARCH}.tar.gz \
    "https://uploads.github.com/repos/${REPO_NAME}/releases/${RELEASE_ID}/assets?name=images-${TAG_NAME}-${OS}-${ARCH}.tar.gz"

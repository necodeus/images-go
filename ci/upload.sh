#!/bin/bash

TOKEN=$1
REPO_NAME=$2
RELEASE_ID=$3
OS=$4
ARCH=$5
VERSION=$6

curl -L \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${TOKEN}" \
    -H "Content-Type: application/octet-stream" \
    --data-binary @./build/images-${OS}-${ARCH}.tar.gz \
    "https://uploads.github.com/repos/${REPO_NAME}/releases/${RELEASE_ID}/assets?name=images-${VERSION}-${OS}-${ARCH}.tar.gz"

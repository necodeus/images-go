#!/bin/bash

TOKEN=$1
REPO_NAME=$2
TAG_NAME=$3

response=$(curl -L \
    -X POST \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${TOKEN}" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    https://api.github.com/repos/${REPO_NAME}/releases \
    -d "{\"tag_name\": \"${TAG_NAME}\", \"name\": \"Release ${TAG_NAME}\", \"draft\": false, \"prerelease\": false}")

release_id=$(echo "$response" | jq -r '.id')

if [ -z "$release_id" ]; then
    release_id=$(curl -L \
        -H "Accept: application/vnd.github+json" \
        -H "Authorization: Bearer ${TOKEN}" \
        https://api.github.com/repos/${REPO_NAME}/releases/tags/${TAG_NAME} | jq -r '.id')
fi

if [ -z "$release_id" ]; then
    echo "Failed to create or get release id"
    exit 1
fi


echo "RELEASE_ID=${release_id}" > release_id

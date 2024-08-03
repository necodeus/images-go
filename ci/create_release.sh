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
    -d "{\"tag_name\": \"${TAG_NAME}\", \"name\": \"${TAG_NAME}\", \"draft\": false, \"prerelease\": false}")

release_id=$(echo "$response" | jq -r '.id')

if [ "$release_id" == "null" ]; then
    response=$(curl -L -H "Accept: application/vnd.github+json" https://api.github.com/repos/${REPO_NAME}/releases/tags/${TAG_NAME})

    release_id=$(echo "$response" | jq -r '.id')
fi

echo "${release_id}" > release_id

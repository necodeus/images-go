#!/bin/bash

OS=$1
ARCH=$2

mkdir -p build
CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -o ./build/images-${OS}-${ARCH}
cp .env.example ./build/.env
cp -r errors ./build/errors
cp init.sql ./build/init.sql

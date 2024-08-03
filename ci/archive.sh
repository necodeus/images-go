#!/bin/bash

OS=$1
ARCH=$2

tar -czvf ./build/images-${ARCH}-${ARCH}.tar.gz -C ./build images-${OS}-${ARCH} .env dot.gif init.sql

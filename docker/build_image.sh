#!/bin/bash

set -e

CURDIR=$(cd "$(dirname "$0")"; pwd)

VER=""
IMAGE_NAME="starcoin-explorer-api"
IMAGE=$IMAGE_NAME
IMAGE_LATEST_TAG=$IMAGE:latest

echo "=== Building linux amd64 binary ==="
cd ../
./build.sh
cp -f ./starcoin-explorer-api_linux_amd64 $CURDIR/image/
cp -rf ./conf $CURDIR/image/

echo "=== Building  image ${IMAGE_LATEST_TAG} ==="
cd  $CURDIR/image
docker build -t $IMAGE_LATEST_TAG .

echo "=== Building done ==="

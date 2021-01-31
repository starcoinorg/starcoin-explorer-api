#!/bin/bash
CURDIR=$(cd "$(dirname "$0")"; pwd)
docker run -d --name starcoin-explorer-api -p 8080:8080  -e "STARCOIN_ES_URL=${STARCOIN_ES_URL}" -e "STARCOIN_ES_USER=${STARCOIN_ES_USER}" -e "STARCOIN_ES_PWD=${STARCOIN_ES_PWD}" starcoin-explorer-api:latest

docker ps



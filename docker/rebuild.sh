#!/bin/bash

./stop.sh
./build_image.sh
./run.sh
docker logs -f starcoin-explorer-api

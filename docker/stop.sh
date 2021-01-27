#!/bin/bash

echo "=== container stopping ==="
docker stop starcoin-explorer-api
docker rm  starcoin-explorer-api
echo "=== container stopped ==="
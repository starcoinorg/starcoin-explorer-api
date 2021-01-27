#!/usr/bin/env bash

GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain --untracked-files=no`" && echo "+CHANGES" || true)
goos=linux
goArch=amd64
BIN_NAME=starcoin-explorer-api_${goos}_${goArch}

echo "==> Building..."
GOOS=$goos GOARCH=$goArch go build -ldflags "-X main.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -o $BIN_NAME
echo "==> Building done"

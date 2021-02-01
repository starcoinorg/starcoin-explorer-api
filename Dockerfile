FROM golang:1.15.7 AS build

ARG GIT_COMMIT
ARG GIT_DIRTY

ENV GIT_COMMIT=$GIT_COMMIT \
    GIT_DIRTY=$GIT_DIRTY \
    goos="linux" \
    goArch="amd64"

ENV BIN_NAME="starcoin-explorer-api_${goos}_${goArch}"

WORKDIR /starcoin-explorer-api
COPY ./ .
RUN go install
RUN GOOS=$goos GOARCH=$goArch go build -ldflags "-X main.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -o $BIN_NAME
RUN ls -la /starcoin-explorer-api

FROM golang:1.15.7

ENV RELEASE_PATH="/starcoin-explorer-api"

WORKDIR /data/starcoin-explorer-api
COPY --from=build /starcoin-explorer-api/starcoin-explorer-api_linux_amd64 \
     /starcoin-explorer-api/docker/entrypoint.sh \
     ./
RUN chmod 755 /data/starcoin-explorer-api/entrypoint.sh

RUN mkdir /data/starcoin-explorer-api/conf
COPY --from=build /starcoin-explorer-api/conf /data/starcoin-explorer-api/conf

RUN ls -la /data/starcoin-explorer-api

ENTRYPOINT ["/data/starcoin-explorer-api/entrypoint.sh"]


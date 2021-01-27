FROM golang:1.15.7
MAINTAINER Ke Wen<wk3368@163.com>

RUN mkdir -p /data/starcoin-explorer-api
COPY starcoin-explorer-api_linux_amd64 /data/starcoin-explorer-api/
COPY conf /data/starcoin-explorer-api/conf
COPY run.sh /
RUN chmod 755 /run.sh


ENTRYPOINT ["/run.sh"]


FROM golang:1.22 AS build 

ARG VERSION
WORKDIR /go/src/accounting

RUN apt-get update \
    && apt-get install -y -q --no-install-recommends 

ENTRYPOINT [ "/bin/sh", "-c", "go test $(go list --buildvcs=false ./...)" ]
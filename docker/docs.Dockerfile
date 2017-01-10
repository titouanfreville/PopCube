FROM golang:1.7-alpine
MAINTAINER Clement LE CORRE <clement@le-corre.eu>

COPY go/src /go/src
WORKDIR /go/src

RUN apk add --no-cache --update git \
    && go get -v golang.org/x/tools/cmd/godoc \
    && rm -rf /go/src/golang.org \
    && apk del git
EXPOSE 6060
ENTRYPOINT godoc -http=:6060 -goroot=/go/src

FROM golang:1.7-alpine

VOLUME /go/

ENTRYPOINT go test -v **/*.go

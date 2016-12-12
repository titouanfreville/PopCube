FROM golang:1.7-alpine

VOLUME /go/
COPY . /go

ENTRYPOINT go test -v **/*.go

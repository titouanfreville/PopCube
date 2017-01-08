FROM golang:1.7-alpine

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

ENV TERM xterm-256color

COPY go/src/api /go/src/api
COPY go/src/models /go/src/models
COPY go/src/utils /go/src/utils
COPY go/src/data_stores /go/src/data_stores
COPY utils/go_get.sh /bin/go_get.sh

RUN apk add --update git bash && \
		cd /go/ && \
		go_get.sh && \
		rm -rf /var/cache/apk/* && \
		rm /bin/go_get.sh

# RUN mv /tmp/go/* /go/ && ls /go && rm -rf /tmp/go
#
ENTRYPOINT go run /go/src/api/api.go

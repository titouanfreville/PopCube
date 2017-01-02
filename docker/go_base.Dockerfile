FROM golang:1.7-alpine

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

RUN ls /go

COPY go/api /go/api
COPY go/models /go/models
COPY go/utils /go/utils
COPY go/data_stores /go/data_stores
COPY utils/go_get.sh /bin/go_get.sh

RUN apk add --update git bash && \
		go_get.sh && \
		rm -rf /var/cache/apk/* && \
		rm /bin/go_get.sh

# RUN mv /tmp/go/* /go/ && ls /go && rm -rf /tmp/go
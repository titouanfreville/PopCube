FROM golang:1.7-alpine

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

COPY go /tmp/go
COPY utils/go_get.sh /bin/go_get.sh
COPY utils/go_test_entrypoint.sh /bin/entrypoint

RUN apk add --update git bash && \
		go_get.sh && \
		rm -rf /var/cache/apk/* && \
		rm /bin/go_get.sh

RUN mv /tmp/go/* /go/ && rm -rf /tmp/go
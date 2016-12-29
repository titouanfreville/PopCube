FROM golang:1.7-alpine

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

COPY api /go/api
COPY models /go/models
COPY data_store /go/data_store
COPY utils/go_get.sh /bin/go_get.sh
COPY utils/go_test_entrypoint.sh /bin/entrypoint

RUN apk add --update git bash && \
		go_get.sh && \
		rm -rf /var/cache/apk/* && \
		rm /bin/go_get.sh
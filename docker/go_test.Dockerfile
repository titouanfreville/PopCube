FROM registry.le-corre.eu:5000/go:base

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com
ENV WATCHING 0

ENTRYPOINT entrypoint /go $WATCHING

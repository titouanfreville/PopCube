FROM registry.le-corre.eu:5000/go:base

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com
ENV WATCHING 0
ENV TERM xterm-256color

ENTRYPOINT entrypoint /go $WATCHING

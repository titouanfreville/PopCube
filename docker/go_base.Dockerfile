FROM golang:1.7-alpine

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

ENV TERM xterm-256color

COPY go/src/api /go/src/api
COPY go/src/models /go/src/models
COPY go/src/utils /go/src/utils
COPY go/src/datastores /go/src/datastores
COPY utils/go_get.sh /bin/go_get.sh

RUN apk add --update git bash && \
		cd /go/ && \
		go_get.sh && \
		rm -rf /var/cache/apk/* && \
		rm /bin/go_get.sh

ENV GOSU_VERSION 1.9
RUN set -x \
    && apk add --no-cache --virtual .gosu-deps \
        dpkg \
        gnupg \
        openssl \
    && dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')" \
    && wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch" \
    && wget -O /usr/local/bin/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch.asc" \
    && export GNUPGHOME="$(mktemp -d)" \
    && gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
    && gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu \
    && rm -r "$GNUPGHOME" /usr/local/bin/gosu.asc \
    && chmod +x /usr/local/bin/gosu \
    && gosu nobody true \
    && apk del .gosu-deps
    
# RUN mv /tmp/go/* /go/ && ls /go && rm -rf /tmp/go
#
ENTRYPOINT go run /go/src/api/api.go

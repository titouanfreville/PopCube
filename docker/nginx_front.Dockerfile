FROM nginx:1.11-alpine
MAINTAINER Clement LE CORRE <clement@le-corre.eu>

ENV TERM xterm-256color

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

RUN apk add --no-cache --update nodejs=6.7.0-r0 \
    && rm -rf /var/cache/apk/*
COPY docs/ /usr/share/nginx/html/docs
COPY app/ /usr/share/nginx/html/app
WORKDIR /usr/share/nginx

RUN npm install -g bower
RUN npm install /usr/share/nginx/html/app
#RUN bower --allow-root install

# This container HAVE TO be run with volume.
FROM ruby:latest
MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

ENV DEBUG 1
ENV WATCHING 0
ENV TERM xterm-256color

RUN apt-get update && apt-get install -qyyy gem-dev ruby-dev
RUN gem install compass --pre

# Gosu installation
ENV GOSU_VERSION 1.9
RUN set -x \
    && apt-get update && apt-get install -y --no-install-recommends ca-certificates wget && rm -rf /var/lib/apt/lists/* \
    && dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')" \
    && wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch" \
    && wget -O /usr/local/bin/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch.asc" \
    && export GNUPGHOME="$(mktemp -d)" \
    && gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
    && gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu \
    && rm -r "$GNUPGHOME" /usr/local/bin/gosu.asc \
    && chmod +x /usr/local/bin/gosu \
    && gosu nobody true \
    && apt-get purge -y --auto-remove ca-certificates wget

RUN mkdir /home/style
VOLUME /home/style
WORKDIR /home/style

COPY utils/sass_compilation.sh /bin/entrypoint
COPY app/styles /home/style

ENTRYPOINT entrypoint scss . $DEBUG $WATCHING
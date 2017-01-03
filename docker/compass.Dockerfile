# This container HAVE TO be run with volume.
FROM ruby:latest

MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

ENV DEBUG 1
ENV WATCHING 0
ENV TERM xterm-256color

RUN apt-get update && apt-get install -qyyy gem-dev ruby-dev
RUN gem install compass --pre

RUN mkdir /home/style
VOLUME /home/style
WORKDIR /home/style

COPY utils/sass_compilation.sh /bin/entrypoint
COPY app/styles /home/style

ENTRYPOINT entrypoint scss . $DEBUG $WATCHING
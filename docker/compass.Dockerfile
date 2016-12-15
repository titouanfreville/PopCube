# This container HAVE TO be run with volume.
FROM ruby:latest

RUN apt-get update && apt-get install -qyyy gem-dev ruby-dev
RUN gem install compass --pre

RUN mkdir /home/style
VOLUME /home/style
WORKDIR /home/style

CMD compass watch --sass-dir scss --force .
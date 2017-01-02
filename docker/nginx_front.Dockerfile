FROM nginx:1.11-alpine
MAINTAINER Clement LE CORRE <clement@le-corre.eu>

RUN apk add --no-cache --update nodejs=6.7.0-r0 \
    && rm -rf /var/cache/apk/*
COPY docs/ /usr/share/nginx/html/docs
COPY app/ /usr/share/nginx/html/app
WORKDIR /usr/share/nginx

RUN npm install -g bower
RUN npm install /usr/share/nginx/html/app
#RUN bower --allow-root install

FROM golang:1.7-alpine


RUN apk add --update git && \
		go get -d -v golang.org/x/tools/cmd/goimports && \
		go get -d -v golang.org/x/tools/cmd/godoc && \
		go get -d -v github.com/golang/lint/golint && \
		go get -d -v github.com/smartystreets/goconvey && \
		go get -d -v golang.org/x/crypto/bcrypt && \
		go get -d -v github.com/nicksnyder/go-i18n/i18n && \
		go get -d -v github.com/pborman/uuid && \
		rm -rf /var/cache/apk/*

VOLUME /go/
COPY ./api /go/api
COPY ./models /go/models

ENTRYPOINT go test -v models/*.go

FROM golang:1.7-alpine


RUN apk add --update git && \
		go get -d -v golang.org/x/tools/cmd/goimports && \
		go get -d -v golang.org/x/tools/cmd/godoc && \
		go get -d -v github.com/golang/lint/golint && \
		go get -d -v github.com/smartystreets/goconvey && \
		go get -d -v golang.org/x/crypto/bcrypt && \
		rm -rf /var/cache/apk/*

VOLUME /go/
COPY . /go


ENTRYPOINT go test -v **/*.go

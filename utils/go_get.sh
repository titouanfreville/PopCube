#!/bin/sh
go get -d -v golang.org/x/tools/cmd/goimports && \
go get -d -v golang.org/x/tools/cmd/godoc && \
go get -d -v github.com/golang/lint/golint && \
go get -d -v github.com/smartystreets/goconvey && \
go get -d -v golang.org/x/crypto/bcrypt && \
go get -d -v github.com/nicksnyder/go-i18n/i18n && \
go get -d -v github.com/pborman/uuid && \
go get -d -v github.com/jinzhu/gorm && \
go get -d -v github.com/Sirupsen/logrus
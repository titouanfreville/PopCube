#!/bin/sh
go get -v golang.org/x/crypto/bcrypt && \
go get -v golang.org/x/tools/cmd/goimports && \
go get -v golang.org/x/tools/cmd/godoc && \
go get -v github.com/golang/lint/golint && \
go get -v github.com/smartystreets/goconvey && \
go get -v github.com/nicksnyder/go-i18n/i18n && \
go get -v github.com/jinzhu/gorm && \
go get -v github.com/pborman/uuid && \
go get -v github.com/Sirupsen/logrus && \
go get -v github.com/alecthomas/log4go && \
go get -v github.com/go-sql-driver/mysql && \
go get -v github.com/pressly/chi
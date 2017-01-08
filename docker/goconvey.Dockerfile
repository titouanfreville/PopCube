FROM registry.le-corre.eu:5000/go:base
MAINTAINER Clement LE CORRE <clement@le-corre.eu>

EXPOSE 8080
CMD ["goconvey","-host","0.0.0.0","-excludedDirs","github.com,golang.org,gopkg.in","-cover"]

FROM registry.le-corre.eu:5000/go:base

# MAINTAINER FREVILLE Titouan titouanfreville@gmail.com

# COPY api /go/api
# COPY models /go/models
# COPY utils/go_get.sh /bin/go_get.sh
# COPY utils/go_test_entrypoint.sh /bin/entrypoint

# RUN apk add --update git bash && \
# 		go_get.sh && \
# 		rm -rf /var/cache/apk/* && \
# 		rm /bin/go_get.sh

ENTRYPOINT entrypoint /go

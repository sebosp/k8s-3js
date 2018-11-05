FROM alpine:3.6
MAINTAINER Seb Osp <kraige@gmail.com>
ENV L0_REFRESHED_AT 20170923
ENV GOROOT="/usr/lib/go"
ENV GOBIN="$GOROOT/bin"
ENV GOPATH="/home/sre/go"
ENV PATH="$PATH:$GOBIN:$GOPATH/bin"
COPY k8s-sniffer.go /app/
WORKDIR /app
RUN apk add --update curl go \
    && go get -v
#ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

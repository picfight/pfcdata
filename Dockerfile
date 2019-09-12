# This Dockerfile is NOT supported.
FROM golang:1.11-alpine3.8
RUN apk update && apk add vim tree lsof bash git gcc musl-dev
ENV GOPATH=/home/decred/go
ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:$GOPATH/bin
ENV PFCSRC_PATH=$GOPATH/src/github.com/picfight/pfcdata/
ENV GO111MODULE=on
RUN adduser -s /bin/bash -D -h /home/decred picfight && chown -R picfight:decred /home/decred
WORKDIR $PFCSRC_PATH
RUN chown -R picfight:decred $GOPATH 
# since we might be rebulding often we need to cache this module layer
# otherwise docker will detect changes everytime and re-download everything again
COPY go.* $PFCSRC_PATH
RUN go mod download 
COPY . $PFCSRC_PATH
RUN chown -R picfight:decred $GOPATH 
USER picfight
RUN go build
CMD /bin/bash

ENTRYPOINT ./pfcdata
# Note: when building the --squash flag is an experimental feature as of Docker 18.06
# docker build --squash -t picfight/pfcdata .
# running
# docker run -ti --rm picfight/pfcdata
# or if attaching source volume and developing interactively
#  docker run -ti --entrypoint=/bin/bash -v ${PWD}:${PWD}:/home/decred/go/src/github.com/picfight/pfcdata --rm picfight/pfcdata
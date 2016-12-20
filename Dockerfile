FROM golang:latest

ENV GOPATH /go

RUN mkdir -p "$GOPATH/src/nchan-auth"
ADD . "$GOPATH/src/nchan-auth"

WORKDIR $GOPATH/src/nchan-auth/app
RUN chmod +x script/*
RUN ./script/build

CMD ./app

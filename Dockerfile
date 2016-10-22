FROM golang:1.7.1-alpine

RUN apk add --update git curl make

COPY . /go/src/github.com/moul/boilergen
WORKDIR /go/src/github.com/moul/boilergen

RUN make install
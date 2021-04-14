FROM golang:1.15.7-alpine AS build
RUN apk update && apk add git
WORKDIR /go/src/app
ADD . /go/src/app

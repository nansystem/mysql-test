FROM golang:1.15.7-alpine AS build
RUN apk update && apk add git
WORKDIR /go/src/app
ADD ./app /go/src/app
RUN go mod download

FROM golang:1.27.3-alpine AS build
RUN apk update && apk add git
WORKDIR /go/src/app
ADD ./app /go/src/app
RUN go mod download

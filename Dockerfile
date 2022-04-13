# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN apk add --no-cache \
        git \
        git-lfs

WORKDIR /app

COPY ./go-server/go.mod ./
COPY ./go-server/go.sum ./

RUN go mod download

COPY ./go-server/*.go ./

RUN go install github.com/camcast3/SimpleCalculatorApi/arithmetic@latest

RUN go build -o /go-server

CMD ["/go-server"]
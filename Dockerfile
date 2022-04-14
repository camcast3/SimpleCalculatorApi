# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN apk add --no-cache \
        git \
        git-lfs

WORKDIR /app

COPY ./go-server/*.mod ./

COPY ./go-server/*.sum ./

COPY ./go-server/*.go ./

RUN mkdir /app/arithmetic

WORKDIR /app/arithmetic

COPY ./go-server/arithmetic/*.go ./

RUN go build .

WORKDIR /app

RUN go mod download

RUN go mod tidy

RUN go build -o /go-server

CMD ["/go-server"]
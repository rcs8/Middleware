FROM golang:1.13-alpine AS build-env
RUN apk update && apk upgrade && apk add --no-cache \
    git \
    build-base \
    gnuplot
ENV GOPATH=/
ENV GO111MODULE=on
WORKDIR /middleware
COPY go.mod .
RUN go mod download
COPY . .
RUN go get -d -v ./...

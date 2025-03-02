# Use official Golang image as the base image
FROM golang:1.23-alpine as builder

RUN apk add --no-cache nginx git curl && apk add busybox-extras
RUN apk upgrade && apk upgrade
RUN apk add --no-cache nginx
RUN apk add --no-cache git
RUN apk add --no-cache curl
RUN apk add --no-cache busybox-extras

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
EXPOSE 5000
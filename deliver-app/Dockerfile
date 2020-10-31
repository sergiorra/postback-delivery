FROM golang:1.15

LABEL maintainer="sergiorra17@gmail.com"
LABEL application="deliver-app"

RUN apt-get update && \
    apt-get upgrade -y

COPY ./deliver-app /go/src/deliverapp
WORKDIR /go/src/deliverapp

RUN go mod tidy


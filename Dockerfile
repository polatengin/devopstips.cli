FROM golang:1.16.2-buster AS build

WORKDIR /src

COPY ./ /src

RUN go build -o devopstips .

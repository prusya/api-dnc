FROM golang:1.19-bullseye as build-env

ADD . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /src/svc

###


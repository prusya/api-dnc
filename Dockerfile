FROM golang:1.19-bullseye as build-env

ADD . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /src/svc

###

FROM gcr.io/distroless/static-debian11

COPY --from=build-env /src/svc /svc

EXPOSE 8080

ENTRYPOINT ["/svc", "run"]
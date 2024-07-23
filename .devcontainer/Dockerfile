FROM golang:1.22.4-alpine3.20

RUN apk update \
    && apk add git \
    && apk add curl

WORKDIR /bin

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xz

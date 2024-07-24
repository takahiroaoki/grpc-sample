FROM golang:1.22.4-alpine3.20

RUN apk update \
    && apk add git \
        curl \
        mysql-client \
        golangci-lint \
        make \
        protobuf

WORKDIR /bin
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xz \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
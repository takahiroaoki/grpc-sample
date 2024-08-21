FROM golang:1.22.4-alpine3.20

RUN apk update \
    && apk add git~=2.45.2 \
        mysql-client~=10.11.8 \
        golangci-lint~=1.57.2 \
        make~=4.4.1 \
        protobuf~=24.4

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && go install github.com/golang/mock/mockgen@v1.6

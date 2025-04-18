FROM amazonlinux:2023.7.20250331.0

# dnf fixes tool versions
RUN dnf install -y \
         git \
         vi \
         make \
         golang \
         protobuf-compiler \
         mariadb105 \
     && rm -rf /var/cache/dnf/* \
     && dnf clean all

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && go install github.com/golang/mock/mockgen@v1.6 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8

ENV PATH $PATH:/root/go/bin

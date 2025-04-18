FROM amazonlinux:2023.7.20250331.0

# dnf fixes tool versions
RUN dnf install -y \
         git \
         make \
         golang \
         protobuf-compiler \
         mariadb105 \
     && rm -rf /var/cache/dnf/* \
     && dnf clean all

#RUN apk update \
#    && apk add \
#        git \
#        mysql-client \
#        golangci-lint~=1.57.2 \
#        make~=4.4.1 \
#        protobuf~=24.4
#
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && go install github.com/golang/mock/mockgen@v1.6

ENV PATH $PATH:/root/go/bin

# grpc-sample

This is the sample grpc server written with golang.

## Requirement
Docker Desktop

※The maintainer use codespaces.

## Tech stack
- Go 1.22
- MySQL 8
- gRPC


Others:
- gorm
- golang-migrate

etc.

## Local setup

```
# migration and generate go code from proto
# the password is necessary for mysql. see .devcontainer/.env
$ make init

# start grpc server
$ make run-server

# start grpc server with reflection
$ make run-server-ref

$ if you need sample data
$ make sample
```

You can try gRPC request via Postman of VSCode extension. The server URL is localhost:8080.

## Appendix
### Unit test
```
$ make test
```

### Lint
```
$ make lint
```

### Create migration files
```
$ go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1
$ migrate create -ext sql -dir /mnt/grpc-sample/app/resource/migration -seq ${MIGRATION_FILE_NAME}
```

### Connect to mysql container from app container
```
# the password is written in .devcontainer/.env
$ make mysql
```

Other commands are written in Makefile.
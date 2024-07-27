# go-env

This is the sample project written with golang.

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
# migrate up, insert data for development, generate go code from proto
# the password is necessary for mysql. see .devcontainer/.env
$ make init

# start grpc server
$ make run-server
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
$ migrate create -ext sql -dir ${PATH_TO_MIGRATION_DIR} -seq ${MIGRATION_FILE_NAME}
```

### Connect to mysql container from app container
```
# the password is written in .devcontainer/.env
$ make mysql
```
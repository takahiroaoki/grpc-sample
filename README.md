# go-env

This is the sample project written with golang.

## Requirement
Docker Desktop

※The maintainer use codespaces.

## Tech stack
- Go 1.22
- MySQL 8


Others:
- gorm
- golang-migrate

etc.

## Local setup

```
# migrate up, insert data for development, generate go code from proto, start grpc server
# the password is necessary for mysql. see .devcontainer/.env
$ make init-run
```

## Appendix
### Lint
```
# lint
$ make lint
```

### How to create migration files
```
$ migrate create -ext sql -dir ${PATH_TO_MIGRATION_DIR} -seq ${MIGRATION_FILE_NAME}
```

### How to connect to mysql container from app container
```
# the password is written in .devcontainer/.env
$ make mysql
```
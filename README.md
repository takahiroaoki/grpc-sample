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
# migrate up
$ make migrate-up

# sample data
$ make insert-dev-data

# start web server, and get access to localhost:8080
$ make run-server
```

## Other commands
```
# lint
$ make lint
```

## Appendix

### Migration on the local environment

#### How to create migration files
```
$ migrate create -ext sql -dir ${PATH_TO_MIGRATION_DIR} -seq ${MIGRATION_FILE_NAME}
```

### How to connect to mysql container from app container
```
# the password is written in .devcontainer/.env
$ make mysql
```
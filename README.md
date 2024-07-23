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
$ migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

# sample data
$ migrate -path "/workspaces/go-env/app/asset/sample" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

# start web server, and get access to localhost:8080
$ cd app
$ go run main.go server
```

## Appendix

### Migration on the local environment

#### How to create migration files
```
$ migrate create -ext sql -dir ${PATH_TO_MIGRATION_DIR} -seq ${MIGRATION_FILE_NAME}
```

#### How to execute migration
```
# up
$ migrate -path "${PATH_TO_MIGRATION_DIR}" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

# down
$ migrate -path "${PATH_TO_MIGRATION_DIR}" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" down

# specific version
$ migrate -path "${PATH_TO_MIGRATION_DIR}" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" goto ${VERSION}
```

### How to connect to mysql container from app container
```
# the password is written in .devcontainer/.env
$ mysql -h demo-mysql -u root -p
```
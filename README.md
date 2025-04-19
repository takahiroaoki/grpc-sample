# grpc-sample

This is the sample grpc server written with golang.

## Requirement
- Docker Desktop
- VSCode & Dev Container Extension

※I use codespaces.

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

## Build

Building binary file can be achieved by ./github/workflows/build.yml via GitHub Actions.

The artifact would be uploaded to [artifact-storage](https://github.com/takahiroaoki/artifact-storage).

That workflow needs personal access token as PERSONAL_ACCESS_TOKEN.

## Deploy

I prepared deploy system for development purpose. See the following repositories.
- [packer-container](https://github.com/takahiroaoki/packer-container)
- [tf-container](https://github.com/takahiroaoki/tf-container)

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
$ migrate create -ext sql -dir /workspaces/grpc-sample/app/resource/migration -seq ${MIGRATION_FILE_NAME}
```

### Connect to mysql container from app container
```
# the password is written in .devcontainer/.env
$ make mysql
```

Other commands are written in Makefile.

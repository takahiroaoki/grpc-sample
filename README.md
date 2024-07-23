# go-env

## Migration on the local environment

### How to create migration files
```
$ migrate create -ext sql -dir /workspaces/go-env/app/asset/migration -seq ${MIGRATION_FILE_NAME}
```

### How to execute migration
```
# up
$ migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

# down
$ migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" down

# specific version
$ migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" goto ${VERSION}
```

## How to connect to mysql container from app container
```
$ mysql -h demo-mysql -u root -p
```
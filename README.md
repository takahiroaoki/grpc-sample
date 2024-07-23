# go-env

## Migration on the local environment

### How to create migration files

Execute the following command.
```
$ migrate create -ext sql -dir /workspaces/go-env/app/asset/migration -seq ${MIGRATION_FILE_NAME}
```

### How to execute migration.

```
$ migrate --path /workspaces/go-env/app/asset/migration --database mysql://root:password@demo-mysql:3306/demodb up
```
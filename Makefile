migrate-up:
	@migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

migrate-down:
	@migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" down

insert-dev-data:
	@migrate -path "/workspaces/go-env/app/asset/dev" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

run-server:
	@cd /workspaces/go-env/app \
	&& go run main.go server

lint:
	@cd /workspaces/go-env/app \
	&& golangci-lint run

mysql:
	@mysql -h demo-mysql -u dev-user -p
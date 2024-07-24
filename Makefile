init-run:migrate-up insert-dev-data proto-go run-server

migrate-up:
	@migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

migrate-down:
	@migrate -path "/workspaces/go-env/app/asset/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" down

insert-dev-data:
	@migrate -path "/workspaces/go-env/app/asset/dev" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

proto-go:
	@protoc --proto_path=proto \
		--go_out=app/pb --go_opt=paths=source_relative \
		--go-grpc_out=app/pb --go-grpc_opt=paths=source_relative \
		sample.proto

run-server:
	@cd /workspaces/go-env/app \
	&& go run main.go server

lint:
	@cd /workspaces/go-env/app \
	&& golangci-lint run

mysql:
	@mysql -h demo-mysql -u dev-user -p
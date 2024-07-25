init:migrate-up insert-dev-data proto-go

migrate-up:
	migrate -path "/workspaces/go-env/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

migrate-down:
	migrate -path "/workspaces/go-env/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" down

insert-dev-data:
	mysql -h demo-mysql -u dev-user -p < /workspaces/go-env/asset/dev.sql

proto-go:
	protoc --proto_path=proto \
		--go_out=app/pb --go_opt=paths=source_relative \
		--go-grpc_out=app/pb --go-grpc_opt=paths=source_relative \
		sample.proto

run-server:
	cd /workspaces/go-env/app \
	&& go run main.go server -p local

lint:
	cd /workspaces/go-env/app \
	&& golangci-lint run

mysql:
	mysql -h demo-mysql -D demodb -u dev-user -p

test:
	cd /workspaces/go-env/app \
	&& go test ./handler ./service ./repository

mockgen:
	rm -f ./app/mock/*_mock.go \
	&& mockgen -source=./app/repository/sample_repository.go -destination=./app/mock/sample_repository_mock.go -package=mock \
	&& mockgen -source=./app/service/sample_service.go -destination=./app/mock/sample_service_mock.go -package=mock
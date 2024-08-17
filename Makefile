# set up
init:db-reset migrate-up proto-go

db-reset:
	mysql -h demo-mysql -u dev-user -p < /workspaces/go-env/devutil/reset.sql

migrate-up:
	migrate -path "/workspaces/go-env/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" up

migrate-down:
	migrate -path "/workspaces/go-env/migration" -database "mysql://root:password@tcp(demo-mysql:3306)/demodb" down

proto-go:
	protoc --proto_path=proto \
		--go_out=app/pb --go_opt=paths=source_relative \
		--go-grpc_out=app/pb --go-grpc_opt=paths=source_relative \
		sample.proto

# run server
run-server:
	cd /workspaces/go-env/app \
	&& go run main.go server

run-server-ref:
	cd /workspaces/go-env/app \
	&& go run main.go server -r true

# test
test:proto-go mockgen
	cd /workspaces/go-env/app \
	&& go clean -testcache \
	&& go test ./handler ./repository ./service

mockgen:
	rm -f ./app/testutil/mock/*_mock.go \
	&& mockgen -source=./app/repository/user_repository.go -destination=./app/testutil/mock/user_repository_mock.go -package=mock \
	&& mockgen -source=./app/service/create_user_service.go -destination=./app/testutil/mock/create_user_service_mock.go -package=mock \
	&& mockgen -source=./app/service/get_user_info_service.go -destination=./app/testutil/mock/get_user_info_service_mock.go -package=mock

# data
db-sample:
	mysql -h demo-mysql -u dev-user -p < /workspaces/go-env/devutil/sample.sql

db-clean:
	mysql -h demo-mysql -u dev-user -p < /workspaces/go-env/devutil/clean.sql

# others
mysql:
	mysql -h demo-mysql -D demodb -u dev-user -p

lint:
	cd /workspaces/go-env/app \
	&& golangci-lint run
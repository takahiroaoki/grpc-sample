# set up
init:proto-go db-reset migrate-up

db-reset:
	mysql -h demo-mysql -u dev-user -D demodb -p < /mnt/grpc-sample/devutil/reset.sql

migrate-up:
	cd ./app \
	&& go run main.go migrate up

migrate-down:
	cd ./app \
	&& go run main.go migrate down

proto-go:
	protoc --proto_path=proto \
		--go_out=app/infra/pb --go_opt=paths=source_relative \
		--go-grpc_out=app/infra/pb --go-grpc_opt=paths=source_relative \
		sample.proto

# run server
run-server:
	cd ./app \
	&& go run main.go server

run-server-ref:
	cd ./app \
	&& go run main.go server -r true

# test
test:proto-go mockgen
	cd ./app \
	&& go clean -testcache \
	&& go test -v ./...

mockgen:
	mockgen -source=./app/domain/service/create_user_service.go -destination=./app/domain/service/create_user_service_mocks.go -package=service
	mockgen -source=./app/domain/service/get_user_info_service.go -destination=./app/domain/service/get_user_info_service_mocks.go -package=service
	mockgen -source=./app/domain/handler/create_user_handler.go -destination=./app/domain/handler/create_user_handler_mocks.go -package=handler
	mockgen -source=./app/domain/handler/get_user_info_handler.go -destination=./app/domain/handler/get_user_info_handler_mocks.go -package=handler
	mockgen -source=./app/infra/server/create_user.go -destination=./app/infra/server/create_user_mocks.go -package=server
	mockgen -source=./app/infra/server/get_user_info.go -destination=./app/infra/server/get_user_info_mocks.go -package=server

# data
db-sample:
	mysql -h demo-mysql -u dev-user -D demodb -p < /mnt/grpc-sample/devutil/sample.sql

db-clean:
	mysql -h demo-mysql -u dev-user -D demodb -p < /mnt/grpc-sample/devutil/clean.sql

# build
build:
	cd app \
	&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc-sample


# others
mysql:
	mysql -h demo-mysql -D demodb -u dev-user -p

lint:
	cd ./app \
	&& golangci-lint run
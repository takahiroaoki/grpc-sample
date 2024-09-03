# set up
init:proto-go db-reset migrate-up

db-reset:
	mysql -h demo-mysql -u dev-user -p < /mnt/grpc-sample/devutil/reset.sql

migrate-up:
	cd ./app \
	&& go run main.go migrate up

migrate-down:
	cd ./app \
	&& go run main.go migrate down

proto-go:
	protoc --proto_path=proto \
		--go_out=app/pb --go_opt=paths=source_relative \
		--go-grpc_out=app/pb --go-grpc_opt=paths=source_relative \
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
	rm -f ./app/testutil/mock/*_mock.go \
	&& mockgen -source=./app/repository/user_repository.go -destination=./app/testutil/mock/user_repository_mock.go -package=mock \
	&& mockgen -source=./app/service/create_user_service.go -destination=./app/testutil/mock/create_user_service_mock.go -package=mock \
	&& mockgen -source=./app/service/get_user_info_service.go -destination=./app/testutil/mock/get_user_info_service_mock.go -package=mock

# data
db-sample:
	mysql -h demo-mysql -u dev-user -p < /mnt/grpc-sample/devutil/sample.sql

db-clean:
	mysql -h demo-mysql -u dev-user -p < /mnt/grpc-sample/devutil/clean.sql

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
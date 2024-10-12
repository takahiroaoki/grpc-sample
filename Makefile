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
	rm -f ./app/testutil/mock/*_mock.go \
	&& mockgen -source=./app/domain/repository/repository_interface.go -destination=./app/testutil/mock/repository_mock.go -package=mock \
	&& mockgen -source=./app/domain/service/service_interface.go -destination=./app/testutil/mock/service_mock.go -package=mock \

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
# set up
init:db-reset migrate-up

db-reset:
	mysql -h demo-mysql -u dev-user -D demodb -p < /workspaces/grpc-sample/devutil/reset.sql

migrate-up:
	git submodule update --init --recursive
	./submodules/artifact-storage/batch-hub/batch-hub-0.1.0 schema-migration up

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
	rm -f ./app/testutil/mockrepository/*.go
	rm -f ./app/testutil/mockservice/*.go
	rm -f ./app/testutil/mockhandler/*.go
	mockgen -source=./app/domain/repository/repository_interface.go -destination=./app/testutil/mockrepository/repository.go -package=mockrepository
	mockgen -source=./app/domain/service/service_interface.go -destination=./app/testutil/mockservice/service.go -package=mockservice
	mockgen -source=./app/domain/handler/handler_interface.go -destination=./app/testutil/mockhandler/handler.go -package=mockhandler

api-test:
	cd ./tests \
	&& go clean -testcache \
	&& go test -v ./...

# data
db-sample:
	mysql -h demo-mysql -u dev-user -D demodb -p < /workspaces/grpc-sample/devutil/sample.sql

db-clean:
	mysql -h demo-mysql -u dev-user -D demodb -p < /workspaces/grpc-sample/devutil/clean.sql

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
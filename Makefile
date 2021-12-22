.PHONY: clean test security build run

APP_NAME = mypoints-rest-api
BUILD_DIR = ./build

start: clean swag
	air

clean:
	rm -rf ./build
	rm -rf ./tmp
	rm -rf ./gen

swag:
	swag init -g app.go

# code below cannot be run at the moment TODO: need to look this up
security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) app.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)
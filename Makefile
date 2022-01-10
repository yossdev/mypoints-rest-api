.PHONY: clean test security build run

APP_NAME = mypoints-rest-api
BUILD_DIR = ./build

start: clean swag test
	air

clean:
	rm -rf ./build
	rm -rf ./tmp
	rm -rf ./gen

swag:
	swag init -g app.go

security:
	gosec -quiet ./...

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) app.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)
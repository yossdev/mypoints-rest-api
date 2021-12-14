.PHONY: clean test security build run

APP_NAME = mypoints-rest-api
BUILD_DIR = ./build

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

swag:
	swag init -g app.go
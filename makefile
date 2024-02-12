.PHONY: default run build test docs clean
# variables
APP_NAME=gopportunities

# tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go -docs
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test -v ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs

# End of file
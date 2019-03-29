.PHONY: all test build

all: clean test build

test:
		go test -v -covermode=count -coverprofile=coverage.out ./...

coverage:
		go test -covermode=count -coverprofile=coverage.out ./...

coverage-html:
		make coverage && go tool cover -html=coverage.out

format:
		go vet ./... && go fmt ./...

build:
		make format && go build -o ./tmp/web-server ./api/main.go

clean:
		rm -r -f ./tmp

lint:
		golangci-lint run

run-dev:
		make format && air -c .air.config

run:
		make api-docs && make format && go run api/main.go

run-cli:
		make format && go run cmd/cli/main.go

api-docs:
		swag init -g api/main.go

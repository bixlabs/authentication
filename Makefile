.PHONY: all test build

all: clean test build

test:
		go test -v ./...

format:
		go vet ./... && go fmt ./...

build:
		make format && go build -o ./tmp/web-server ./cmd/api/main.go

clean:
		rm -r -f ./tmp

lint:
		golangci-lint run

run-dev:
		make format && ~/.air -c .air.config

run:
		make format && go run cmd/api/main.go

deps:
		sh ./scripts/install_dep.sh
		sh ./scripts/install_air.sh
		sh ./scripts/install_golangci_lint.sh
		dep ensure

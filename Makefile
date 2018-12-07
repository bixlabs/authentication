.PHONY: all test build

all: clean test build

test:
		go test -v ./...

build:
		go build -o ./tmp/web-server ./cmd/api/main.go

clean:
		rm -r -f ./tmp

run-dev:
		~/.air -c .air.config

run:
		$ go run cmd/api/main.go

deps:
		sh ./scripts/install_dep.sh
		sh ./scripts/install_air.sh
		dep ensure

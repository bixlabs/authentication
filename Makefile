all: deps lint

.PHONY: test clean format lint coverage coverage-html build build-for-mac build-for-windows

deps:
		./deps.sh

test:
		go test -v ./...

coverage:
		go test -covermode=count -coverprofile=coverage.out ./...

coverage-html:
		make coverage && go tool cover -html=coverage.out

format:
		go vet ./... && go fmt ./...

build:
		make api-docs && make format && go build --tags "sqlite_userauth" -o ./tmp/auth-server ./api/main.go

build-for-mac:
		GOOS=darwin GOARCH=amd64 make build

build-for-windows:
		GOOS=windows GOARCH=386 make api-docs && make format && go build -o ./tmp/auth-server.exe ./api/main.go

build-for-docker:
		make build

clean:
		rm -r -f ./tmp

lint:
		$(GOPATH)/bin/golangci-lint run --enable-all --disable goimports

run-dev:
		make format && $(GOPATH)/air -c .air.config

run:
		make api-docs && make format && go run api/main.go

run-cli:
		make format && go run cmd/cli/main.go

api-docs:
		$(GOPATH)/bin/swag init -g ./api/main.go -o ./api/docs

ci:
		make all build

docker-build:
		docker-compose -f deployments/docker-compose.yml build

docker-run:
		docker-compose -f deployments/docker-compose.yml up -d

docker-run-dev:
		docker-compose -f deployments/docker-compose.yml -f deployments/docker-compose.dev.yml up

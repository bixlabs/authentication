#!/usr/bin/env bash

echo "package docs" > ./api/docs/docs.go
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
go get -u github.com/swaggo/swag/cmd/swag
make test

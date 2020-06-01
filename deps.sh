#!/usr/bin/env sh

mkdir -p ./api/docs
echo "package docs" > ./api/docs/docs.go
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
go get -u github.com/swaggo/swag/cmd/swag

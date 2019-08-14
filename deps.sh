#!/usr/bin/env bash

curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.17.1
go get -u github.com/swaggo/swag/cmd/swag
make test

curl -fLo ${GOPATH}/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
chmod +x ${GOPATH}/air

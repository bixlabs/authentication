#!/usr/bin/env bash

# We have to do some witchcraft here because of https://github.com/ugorji/go/issues/279, otherwise we obtain an ambiguous import problem.
go get github.com/ugorji/go@v1.1.2-0.20180831062425-e253f1f20942
make test
go get github.com/ugorji/go@v1.1.2-0.20180831062425-e253f1f20942
go get -u github.com/swaggo/swag/cmd/swag
go get github.com/ugorji/go@v1.1.2-0.20180831062425-e253f1f20942
# We install master version because of this: https://github.com/golangci/golangci-lint/issues/479, should be changed later to the latest tag.
GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@master

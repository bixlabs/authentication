#!/usr/bin/env bash

go get -u github.com/swaggo/swag/cmd/swag
source $GOPATH/bin
make api-docs
make test

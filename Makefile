all: deps lint

.PHONY: test clean format lint coverage coverage-html build

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

build: build-linux build-for-mac build-for-windows

build-linux: format api-docs
		go build --tags "sqlite_userauth" -o ./tmp/authenticator-server-linux ./api/main.go

build-for-mac: format api-docs
		CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build --tags "sqlite_userauth" -o ./tmp/authenticator-server-mac ./api/main.go

build-for-windows: format api-docs
		GOOS=windows GOARCH=386 go build -o ./tmp/authenticator-server-win.exe ./api/main.go

build-admin-cli:
		make format && go build -o ./tmp/admincli ./admincli/main.go

clean:
		rm -r -f ./tmp

lint:
		./bin/golangci-lint run --enable-all \
			-D goimports \
			-D godox \
			-D wsl \
			-D godot \
			-D goerr113 \
			--timeout 2m0s

run:
		make api-docs && make format && go run api/main.go

api-docs:
		$(GOPATH)/bin/swag init --generalInfo  ./api/main.go --output ./api/docs

cli-docs:
		go run ./admincli/docs/main.go
ci:
		make all build-linux build-windows

docker-build:
		docker-compose -f deployments/docker-compose.yml build

docker-run:
		docker-compose -f deployments/docker-compose.yml up -d

		
#####################################################
## Deployments
#####################################################
terraform-aws-root = deployments/terraform/aws

# .tfvars file path for terraform deployment
terraform-aws-tfvars = $(terraform-aws-root)/$(ENV).aws.tfvars

define terraform_aws_apply
	terraform init $(terraform-aws-root) && terraform apply -var-file="$(terraform-aws-tfvars)" $(terraform-aws-root)
endef

deploy-aws:
		[ -f $(terraform-aws-tfvars) ] && $(call terraform_aws_apply) || echo "Error: $(ENV).aws.tfvars does not exists"
		


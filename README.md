# Requirements

* go 1.11 or higher.
* Make sure that you activate [Go Modules](https://github.com/golang/go/wiki/Modules#how-to-install-and-activate-module-support)

## Installation, please read this before running anything.

```bash
$ make
```

* Through this we install some binaries and all the go libraries that the project needs.
* We also run the tests and the linter.

## Go Modules
* If you need more information about them please go [here](https://github.com/golang/go/wiki/Modules#how-to-define-a-module)
* To install new dependencies just use `go get`

## Creating an API through [gin-gonic](https://github.com/gin-gonic/gin)

```bash
$ make run
```

* this will run an HTTP server in port 8080, if you want to change this port you have to specify it through the environment variable (AUTH_SERVER_PORT) in the _.env_ file.
* For testing all the define endpoints you can try out these different CURL commands:
```bash
    * Create: $ curl -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:8080/todo"
    * Read: $ curl -X GET "http://localhost:8080/todo/1"
    * Update: $ curl -X PUT -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:8080/todo"
    * Delete: $ curl -X DELETE "http://localhost:8080/todo/1"
```

## How to generate API documentation
* Initialize the documentation (this will generate a _docs_ folder in the root folder)
```bash
$ make api-docs
```
or
```bash
$ swag init -g ./api/main.go
```
* Then we need to run the project as a web server and go to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html#)

    
## Running the project to show output in console

* Because we are using Clean Architecture, we want to show how the same code is running in different ways without too much effort:
```bash
 $ make run-cli
```
* You should see the output of the CRUD in your console.

## How to make build of a main.go file and run it on Linux.

```bash
$ make build
```
or 

```bash
$ go build -o ./tmp/auth-server ./api/main.go
```
* The command above will create a file called `auth-server` in folder _tmp_, that file is an executable with the main in _./api/main.go_
* To run your executable you have to:
    * Make it executable: `chmod +x ./tmp/auth-server`
    * Run it: `./tmp/auth-server`
* You can build whatever you want (it doesn't have to be a web-server), for example there is another main with which you can follow the same steps _./cmd/cli/main.go_

## Building for MacOS
```bash
$ make build-for-mac
```
* Same steps as above, an executable will be created in _./tpm/_

## Building for Windows
```bash
$ make build-for-windows
```
* Same steps as above, an executable will be created in _./tpm/_

## Hot reload for the Web Server

* To run the project with hot reload: 
```bash
$ make run-dev
```
or 

```bash
$ air -c .air.config
```
* IMPORTANT: Hot reload will not work in MacOS or windows, you have to change the configuration of `.air.config` to instead of using `make build` to use the correct build process depending on the OS you are. For windows you also have to change the _bin_ to have the _.exe_ extension.


## How to run format

* We use `go vet` and `go fmt` for simple linter and formatting, running `make format` will do.
* This commands are run also when we run the project, you can check the [Makefile](./Makefile) to know exactly in which commands is used.

## How to run the linter

* For the linter we are using [golangci-lint](https://github.com/golangci/golangci-lint)
* To run it you can either use:
```bash
$ make lint
```
or 

```bash
$ golangci-lint run
```

## Logging framework

* We are using [Logrus](https://github.com/sirupsen/logrus) as a logging framework
* This is how we initialize the logger [here](./tools/logger.go), specifically `InitializeLogger`
* We have to run `InitializeLogger` before using the `Log` function, [here's an example](./cmd/api/main.go)
* Finally [examples of using the logger](todo/interactors/todo_handler.go)
* visit their website for advance information on how to use it.
* When using `make run-dev` we won't see the colors of the log message, with `make run` we will.

## How to handle environment variables

* For environment variables we use the same `.env` mechanism that we all know, for more information here's the [library](https://github.com/joho/godotenv)
* You can either use the mechanism to read the environment variables from the `.env` file that's explain in the library above OR use this one in this [other library](https://github.com/caarlos0/env)
* [Here's how we load](./cmd/api/main.go#L8) the `.env` file
* For an example you can check [here](./api/todo.go#L15), we are using the second library method which let us use tags in structs. Then we _load_ the struct with the values like [this](./api/todo.go#L21)
* For testing this you can create a `.env` file with a different port than the default, you will see how the web server is initialize in the port you specified in the `.env`, you can just change the name of `.env-template` to `.env` and that will do the trick.

## How to run the tests

```bash
$ make test
```
or 

```bash
$ go test -cover -v ./...
```
* this will run all the test files in the project.
* Test should be in the same folder of the file they are testing and the file name of the test must have the suffix `_test`, if you see the example in _test_ folder you will get it right away.
* Gomega is being used for improving assertion mechanism.

## How to see test coverage

* will show information about the coverage: 
```bash
$ make test
```
* if you want to only see the test coverage information (without the tests logs):


```bash
$ make coverage
```
* if you want to see the coverage in detail in a browser:

```bash
$ make coverage-html
```
* For knowing how are we generating test coverage please check [this](https://blog.golang.org/cover)

# Storage

* This project uses an SQLite database as default storage, for now this will be the only storage capability of the project.
Possibly in the future we will have different ways of providing external storage to the system.

* We are also using an [ORM](https://gorm.io/) library.

## TL;DR How to run/build

* Build:  `$ make build` or `$ go build -o <destination_of_executable_relative_to_root> <path_of_main_file_relative_to_root>`
* Run:
    * Without executable: `$ make run` or `$ go run <path_of_main_file_relative_to_root>`
    * With executable:
        * Make the file executable: `$ chmod +x <path_to_executable_relative_to_root>`
        * Run it: `$ <path_to_executable_relative_to_root>`

## Badges

* [Go Report Card](https://goreportcard.com/) - It will scan your code with `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Replace `github.com/bixlabs/go-layout` with your project reference.

* [GoDoc](http://godoc.org) - It will provide online version of your GoDoc generated documentation. Change the link to point to your project.

* Release - It will show the latest release number for your project. Change the github link to point to your project.

[![Go Report Card](https://goreportcard.com/badge/github.com/bixlabs/go-layout?style=flat-square)](https://goreportcard.com/report/github.com/bixlabs/go-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Docker

### Getting Started

At the time of configuration we used Docker version 19.03.1 and Docker Compose version 1.24.1

#### Prerequisites

In order to run this container you'll need [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/) installed.
You will also need to configure the `.env` in the root of the project, please read [the environment section](#how-to-handle-environment-variables).

#### Usage

To run the container you should run:

```bash
$ make docker-run
```

If you need to build the image:

```bash
$ make docker-build
```

For development you must run the project with hot reload:

```bash
$ make docker-run-dev
```

Behind the scenes we are using docker-compose, 
but you could build and run the same configuration using docker:

```bash
$ docker build -f build/package/Dockerfile -t authentication_api .
```

```bash
$ docker run -d --env-file=.env -p 9000:9000 --name authentication_api authentication_api
```

Also with hot reload:

```bash
$ docker run -d --env-file=.env -p 9000:9000 --name authentication_api authentication_api make run-dev
```

## Notes

If you want more information about all the folders being used in this project please refer to the [original template](https://github.com/golang-standards/project-layout). Thanks for the authors of this one!

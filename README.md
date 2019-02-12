# Requirements

* go 1.11 or higher.

## Go Modules
* If you need more information about them please go [here](https://github.com/golang/go/wiki/Modules#how-to-define-a-module)
* To install dependencies just use `go get`

## Creating an API through [gin-gonic](https://github.com/gin-gonic/gin)

```bash
$ make run
```
or 

```bash
$ go run cmd/api/main.go
```
* this will run an HTTP server in port 3000
* For testing all the define endpoints you can try out these different CURL commands:
```bash
    * Create: $ curl -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:3000/todo"
    * Read: $ curl -X GET "http://localhost:3000/todo/1"
    * Update: $ curl -X PUT -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:3000/todo"
    * Delete: $ curl -X DELETE "http://localhost:3000/todo/1"
```
    
## Running the project to show output in console

* Because we are using Clean Architecture, we want to show how the same code is running in different ways without too much effort:
```bash
 $ make run-cli
```
* You should see the output of the CRUD in your console.

## How to make build of a main.go file and run it.

```bash
$ make build
```
or 

```bash
$ go build -o ./tmp/web-server ./cmd/api/main.go
```
* The command above will create a file called `web-server` in folder _tmp_, that file is an executable with the main in _./cmd/api/main.go_
* To run your executable you have to:
    * Make it executable: `chmod +x ./tmp/web-server`
    * Run it: `./tmp/web-server`
* You can build whatever you want (it doesn't have to be a web-server), for example there is another main with which you can follow the same steps _./cmd/cli/main.go_

## Hot reload for the Web Server

* To run the project with hot reload: 
```bash
$ make run-dev
```
or 

```bash
$ air -c .air.config
```


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
* Finally [examples of using the logger](./todo/use_cases/todo_handler.go)
* visit their website for advance information on how to use it.
* When using `make run-dev` we won't see the colors of the log message, with `make run` we will.

## How to handle environment variables

* For environment variables we use the same `.env` mechanism that we all know, for more information here's the [library](https://github.com/joho/godotenv)
* You can either use the mechanism to read the environment variables from the `.env` file that's explain in the library above OR use this one in this [other library](https://github.com/caarlos0/env)
* [Here's how we load](./cmd/api/main.go#L6) the `.env` file
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

## Notes

If you want more information about all the folders being used in this project please refer to the [original template](https://github.com/golang-standards/project-layout). Thanks for the authors of this one!

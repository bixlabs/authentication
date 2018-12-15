# Requirements
* go 1.9 or higher.

# Installation & Usage

* First of all we have to ensure that our project is in the correct path, follow option #1 in this [dep tutorial](https://golang.github.io/dep/docs/new-project.html) to know where your project should be.
* Run `$ make deps` to install all the dependencies.

## Using Dep

* `dep ensure` ensures that all the go dependencies are installed (this is run in the Makefile).
* If you are going from scratch and don't need the examples provided in this template you will need to remove `Gopkg.lock` and `Gopkg.toml` and initialize dep with `$ dep init`
* After running `$ dep init` a _vendor_ folder will be created and you can _add your dependencies_ with the following `$ dep ensure -add github.com/foo/bar github.com/baz/quux` if you want to know more about dep just check the tutorial above.
* To understand how are you going to be using dep please check [this](https://golang.github.io/dep/docs/daily-dep.html#using-dep-ensure) out

## Creating an API through [gin-gonic](https://github.com/gin-gonic/gin)

* When installing the dependencies for this project you will get gin-gonic for API definition.
* To run the main for the API application you have to `$ make run` or `$ go run cmd/api/main.go`, this will run an HTTP server in port 8080
* For testing all the define endpoint you can try out these different CURL commands:
    * Create: `$ curl -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:8080/todo"`
    * Read: `$ curl -X GET "http://localhost:8080/todo/1"`
    * Update: `$ curl -X PUT -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:8080/todo"`
    * Delete: `$ curl -X DELETE "http://localhost:8080/todo/1"`

## How to make build of a main.go file and run it.
* `$ make build` or `$ go build -o ./tmp/web-server ./cmd/api/main.go`
* The command above will create a file called `web-server` in folder _tmp_, that file is an executable with the main in _./cmd/api/main.go_
* To run your executable you have to:
    * Make it executable: `chmod +x ./tmp/web-server`
    * Run it: `./tmp/web-server`
* You can build whatever you want (it doesn't have to be a web-server), for example there is another main with which you can follow the same steps _./cmd/cli/main.go_

## Hot reload for the Web Server
* For Mac/Windows users you will have to modify [install_dep.sh](./scripts/install_dep.sh) and change the OS variable to match your Operating System.
* We have to install [Air](https://github.com/cosmtrek/air) (if you ran `$ make deps` command this should be already installed and you can skip this):
    * `curl -fLo ~/.air \
           https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air`
    * `chmod +x ~/.air`
* To run the project with hot reload: `$ make run-dev` or `~/.air -c .air.config`

## How to run the tests
* After you install everything with dep you should be able to do `$ make test` or `$ go test -cover -v ./...` this will run all the test files in the project.
* Test should be in the same folder of the file they are testing and the file name of the test must have the suffix `_test`, if you see the example in _test_ folder you will get it right away.

## TL;DR How to run/build
* Build:  `$ make build` or `$ go build -o <destination_of_executable_relative_to_root> <path_of_main_file_relative_to_root>`
* Run:
    * Without executable: `$ make run` or `$ go run <path_of_main_file_relative_to_root>`
    * With executable:
        * Make the file executable: `$ chmod +x <path_to_executable_relative_to_root>
        * Run it: `$ <path_to_executable_relative_to_root>

# Standard Go Project Layout

This is a basic layout for Go application projects. It's not an official standard defined by the core Go dev team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem. Some of these patterns are more popular than others. It also has a number of small enhancements along with several supporting directories common to any large enough real world application.

If you are trying to learn Go or if you are building a PoC or a toy project for yourself this project layout is an overkill. Start with something really simple (a single `main.go` file is more than enough). As your project grows keep in mind that it'll be important to make sure your code is well structured otherwise you'll end up with a messy code with lots of hidden dependencies and global state. When you have more people working on the project you'll need even more structure. That's when it's important to introduce a common way to manage packages/libraries. When you have an open source project or when you know other projects import the code from your project repository that's when it's important to have private packages and code. Clone the repository, keep what you need and delete everything else! Just because it's there it doesn't mean you have to use it all. None of these patterns are used in every single project. Even the `vendor` pattern is not universal.

This project layout is intentionally generic and it doesn't try to impose a specific Go package structure.

This is a community effort. Open an issue if you see a new pattern or if you think one of the existing patterns needs to be updated.

If you need help with naming, formatting and style start by running [`gofmt`](https://golang.org/cmd/gofmt/) and [`golint`](https://github.com/golang/lint). Also make sure to read these Go code style guidelines and recommendations:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments

See [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) for additional background information.

More about naming and organizing packages as well as other code structure recommendations:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

## Go Directories

### `/todo`

This directory varies according to the name of the application, in this example it is _todo_ because that's what this example is about. You will have all
your business related code under this directory.

### `/cmd`

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

See the [`/cmd`](cmd/README.md) directory for examples.

### `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool like [`dep`](https://github.com/golang/dep)).

Don't commit your application dependencies if you are building a library.

## Service Application Directories

### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

## Common Application Directories

### `/configs`

Configuration file templates or default configs.

Put your `confd` or `consul-template` template files here.

### `/scripts`

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple (e.g., `https://github.com/hashicorp/terraform/blob/master/Makefile`).

See the [`/scripts`](scripts/README.md) directory for examples.

### `/build`

Packaging and Continuous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.

Put your CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the `/build/ci` directory linking them to the location where the CI tools expect them (when possible).

### `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).

### `/test`

Additional external test apps and test data. Feel free to structure the `/test` directory anyway you want. For bigger projects it makes sense to have a data subdirectory. For example, you can have `/test/data` or `/test/testdata` if you need Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "_", so you have more flexibility in terms of how you name your test data directory.

See the [`/test`](test/README.md) directory for examples.

## Other Directories

### `/tools`

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.

See the [`/tools`](tools/README.md) directory for examples.

## Directories You Shouldn't Have

### `/src`

Some Go projects do have a `src` folder, but it usually happens when the devs came from the Java world where it's a common pattern. If you can help yourself try not to adopt this Java pattern. You really don't want your Go code or Go projects to look like Java :-)

Don't confuse the project level `/src` directory with the `/src` directory Go uses for its workspaces as described in [`How to Write Go Code`](https://golang.org/doc/code.html). The `$GOPATH` environment variable points to your (current) workspace (by default it points to `$HOME/go` on non-windows systems). This workspace includes the top level `/pkg`, `/bin` and `/src` directories. Your actual project ends up being a sub-directory under `/src`, so if you have the `/src` directory in your project the project path will look like this: `/some/path/to/workspace/src/your_project/src/your_code.go`. Note that with Go 1.11 it's possible to have your project outside of your `GOPATH`, but it still doesn't mean it's a good idea to use this layout pattern.


## Badges

* [Go Report Card](https://goreportcard.com/) - It will scan your code with `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Replace `github.com/golang-standards/project-layout` with your project reference.

* [GoDoc](http://godoc.org) - It will provide online version of your GoDoc generated documentation. Change the link to point to your project.

* Release - It will show the latest release number for your project. Change the github link to point to your project.

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notes

A more opinionated project template with sample/reusable configs, scripts and code is a WIP.

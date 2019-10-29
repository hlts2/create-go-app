# create-go-app
create-go-app is the command like [create-react-app](https://github.com/facebook/create-react-app)

## Requirement
Go (>=1.11)

## Installation
```shell
go get github.com/hlts2/create-go-app/cmd/create-go-app
```

## Usage

```
$ create-go-app --help

Generate skeleton for a Go project

Usage:
  create-go-app [command]

Examples:

* Basic usage:
      create-go-app init golang_sample --mod github.com/user/golang_sample


* Generate CLI skeleton:
      create-go-app cli-init golang_cli_sample --mod github.com/user/golang_cli_sample


Available Commands:
  cli-init    Initialize the Go CLI application
  help        Help about any command
  init        Initialize the Go application

Flags:
  -h, --help   help for create-go-app

Use "create-go-app [command] --help" for more information about a command.
```

### Initialize a new project

```shell
$ create-go-app init golang_sample --mod github.com/user/golang_sample
✔ README.md
✔ go.mod
✔ config/config.yaml
✔ main.go
✔ pkg/config/config.go
✔ pkg/infra/repository/api/repository.go
✔ pkg/infra/repository/db/repository.go
✔ pkg/domain/service/service.go
✔ pkg/domain/model/model.go
✔ pkg/usecase/usecase.go

Scaffold golang-sample successfully 🎉
▸ You can run application with `cd golang_sample; go run main.go -c config/config.yaml` command

$ cd golang_sample
$ tree .
.
├── README.md
├── config
│   └── config.yaml
├── go.mod
├── main.go
└── pkg
    ├── config
    │   └── config.go
    ├── domain
    │   ├── model
    │   │   └── model.go
    │   └── service
    │       └── service.go
    ├── infra
    │   └── repository
    │       ├── api
    │       │   └── repository.go
    │       └── db
    │           └── repository.go
    └── usecase
        └── usecase.go
```

### TODO

- [ ] test code
- [ ] CLI Generator

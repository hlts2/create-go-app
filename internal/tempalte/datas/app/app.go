package app

var ReadmeTmpl = `
# {{.AppName}}
`

var ModTmpl = `
module {{.Mod}}

go 1.12
`

var ConfigYAMLTmpl = `
debug: true
`

var MainTmpl = `
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kpango/glg"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"

	"{{.Mod}}/pkg/config"
	"{{.Mod}}/pkg/usecase"
)

type params struct {
	configPath string
}

func parseParams() *params {
	p := new(params)

	pflag.StringVarP(&p.configPath,
		"config",
		"c",
		"config.yaml",
		"sets configuration file",
	)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	return p
}

func run(params *params) error {
	cfg, err := config.Load(params.configPath)
	if err != nil {
		return errors.Wrap(err, "faild to load configurationg")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := usecase.New{{.AppName|ToCamel}}Usecase(cfg).Start(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	defer func() {
		signal.Stop(sigCh)
		close(sigCh)
		close(errCh)
	}()

	for {
		select {

		// ユースケース層のエラー
		case err := <-errCh:
			return errors.WithStack(err)

		// シグナルトラップ
		case sig := <-sigCh:
			glg.Infof("received os signal: %v", sig)
			cancel()
		}
	}
}

func main() {
	if err := run(parseParams()); err != nil {
		fmt.Fprintf(os.Stderr, "The application has terminated, because an error occurred: %v", err)
		os.Exit(1)
	}
}
`

var ConfigTmpl = `
package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	yaml "gopkg.in/yaml.v2"
)

type Config struct{
	Debug string
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "faild to open configuration file: %v", path)
	}
	defer f.Close()

	var cfg Config

	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "faild to decode")
	}

	if err := validator.New().Struct(cfg); err != nil {
		return nil, errors.Wrap(err, "invalid config object")
	}

	return &cfg, nil
}
`

var APIRepositoryTmpl = `
package api

type {{.AppName|ToCamel}}Repository interface{}

type {{.AppName|ToLowerCamel}}Repository struct{}

func New{{.AppName|ToCamel}}Repository() {{.AppName|ToCamel}}Repository {
	return new({{.AppName|ToLowerCamel}}Repository)
}
`

var DBRepositoryTmpl = `
package db

type {{.AppName|ToCamel}}Repository interface{}

type {{.AppName|ToLowerCamel}}Repository struct{}

func New{{.AppName|ToCamel}}Repository() {{.AppName|ToCamel}}Repository {
	return new({{.AppName|ToLowerCamel}}Repository)
}
`

var ServiceTmpl = `
package service

import (
	"{{.Mod}}/pkg/infra/repository/api"
	"{{.Mod}}/pkg/infra/repository/db"
)

type {{.AppName|ToCamel}}Service interface{}

type {{.AppName|ToLowerCamel}}Service struct {
	apiRepository api.{{.AppName|ToCamel}}Repository
	dbRepository  db.{{.AppName|ToCamel}}Repository
}

func New{{.AppName|ToCamel}}Service(
	apiRepository api.{{.AppName|ToCamel}}Repository,
	dbRepository db.{{.AppName|ToCamel}}Repository,
) {{.AppName|ToCamel}}Service {
	return &{{.AppName|ToLowerCamel}}Service{
		apiRepository: apiRepository,
		dbRepository:  dbRepository,
	}
}
`

var ModelTmpl = `
package model

type {{.AppName|ToCamel}}Model struct {}
`

var UsecaseTmpl = `
package usecase

import (
	"context"

	"{{.Mod}}/pkg/config"
	"{{.Mod}}/pkg/infra/repository/api"
	"{{.Mod}}/pkg/infra/repository/db"
	"{{.Mod}}/pkg/domain/service"
)

type {{.AppName|ToCamel}}Usecase interface {
	Start(ctx context.Context) chan error
}

type {{.AppName|ToLowerCamel}}Usecase struct {
	{{.AppName|ToLowerCamel}}Service service.{{.AppName|ToCamel}}Service
}

func New{{.AppName|ToCamel}}Usecase(cfg *config.Config) {{.AppName|ToCamel}}Usecase {
	return &{{.AppName|ToLowerCamel}}Usecase{
		{{.AppName|ToLowerCamel}}Service: service.New{{.AppName|ToCamel}}Service(
			api.New{{.AppName|ToCamel}}Repository(),
			db.New{{.AppName|ToCamel}}Repository(),
		),
	}
}

func (d *{{.AppName|ToLowerCamel}}Usecase) Start(ctx context.Context) chan error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- d.start(ctx)
	}()

	return errCh
}

func (d *{{.AppName|ToLowerCamel}}Usecase) start(ctx context.Context) error {
	return nil
}
`

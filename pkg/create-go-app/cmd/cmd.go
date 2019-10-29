package cmd

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/hlts2/create-go-app/internal/generator"
)

type params struct {
	mod string
}

// NewCommand creates a create command.
func NewCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-go-app",
		Short: "Generate skeleton for a Go project",
		Example: `
* Basic usage:
      create-go-app init golang_sample --mod github.com/user/golang_sample


* Generate CLI skeleton:
      create-go-app cli-init golang_cli_sample --mod github.com/user/golang_cli_sample
`,
	}

	// register sub command of `create-go-app` command.
	cmd.AddCommand(
		newInitCommand(ctx),
		newInitCLICommand(ctx),
	)

	return cmd
}

func newInitCommand(ctx context.Context) *cobra.Command {
	p := new(params)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the Go application",
		Example: `
* Basic usage:
      create-go-app init golang_sample --mod github.com/user/golang_sample
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			gp, err := getGeneratorParams(args, p)
			if err != nil {
				return errors.WithStack(err)
			}

			g := generator.NewGenerator(gp, cmd.OutOrStdout())
			g.Generate()

			return nil
		},
	}

	// sets command options.
	cmd.Flags().StringVar(&p.mod, "mod", "", "sets go module name")

	return cmd
}

func newInitCLICommand(ctx context.Context) *cobra.Command {
	p := new(params)

	cmd := &cobra.Command{
		Use:   "cli-init",
		Short: "Initialize the Go CLI application",
		Example: `
* Basic usage:
      create-go-app cli-init golang_sample --mod github.com/user/golang_sample
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			gp, err := getGeneratorParams(args, p)
			if err != nil {
				return errors.WithStack(err)
			}

			g := generator.NewGenerator(gp, cmd.OutOrStdout(), generator.WithCLI())
			g.Generate()

			return nil
		},
	}

	// sets command options.
	cmd.Flags().StringVar(&p.mod, "mod", "", "sets go module name")

	return cmd
}

func getGeneratorParams(args []string, params *params) (*generator.Params, error) {
	if len(args) == 0 || len(args[0]) == 0 {
		return nil, errors.New("application name is empty")
	}
	appName := args[0]

	mod := params.mod
	if len(params.mod) == 0 {
		mod = appName
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "faild to get current directory")
	}

	return &generator.Params{
		Wd:      wd,
		AppName: appName,
		Mod:     mod,
	}, nil
}

package generator

import (
	apptemplate "github.com/hlts2/create-go-app/internal/tempalte"
)

// Option is functional option for generator object.
type Option func(g *generator)

// WithCLI rturns option that sets CLI mode.
func WithCLI() Option {
	return func(g *generator) {
		g.templates = apptemplate.CLIApplicationTemplates()
	}
}

package generator

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"

	apptemplate "github.com/hlts2/create-go-app/internal/tempalte"
)

// Generator is an interface to create go source code.
type Generator interface {
	Generate()
}

// Prams is parameter for generated package.
type Params struct {
	Wd      string
	AppName string
	Mod     string
}

type generator struct {
	params    *Params
	outW      io.Writer
	bytePool  sync.Pool
	templates []apptemplate.Templater
}

// NewGenerator returns Generator implementations(*generator).
func NewGenerator(params *Params, outW io.Writer, opts ...Option) Generator {
	return (&generator{
		params: params,
		outW:   outW,
		bytePool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
		templates: apptemplate.ApplicationTemplates(),
	}).apply(opts...)
}

func (g *generator) apply(opts ...Option) *generator {
	for _, op := range opts {
		op(g)
	}
	return g
}

func (g *generator) Generate() {
	ok := true

	for _, t := range g.templates {
		err := g.generate(t.Path(), t.Template(), g.params)
		if err != nil {
			ok = false
			fmt.Fprintf(g.outW, "%4s %s\n", aurora.Red("×"), t.Path())
		} else {
			fmt.Fprintf(g.outW, "%4s %s\n", aurora.Green("✔"), t.Path())
		}
	}

	if !ok {
		return
	}

	fmt.Fprintln(g.outW)
	fmt.Fprintln(g.outW, "Scaffold", aurora.BrightWhite(strcase.ToKebab(g.params.AppName)), "successfully 🎉")
	fmt.Fprintln(g.outW, aurora.Yellow("▸"), fmt.Sprintf("You can run application with `cd %s; go run main.go -c config/config.yaml` command", g.params.AppName))
}

func (g *generator) generate(path string, t *template.Template, params *Params) error {
	buf := g.bytePool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		g.bytePool.Put(buf)
	}()

	if err := t.Execute(buf, params); err != nil {
		return errors.WithStack(err)
	}

	absPath := filepath.Join(g.params.Wd, params.AppName, path)

	b, err := goFmt(absPath, buf.Bytes())
	if err != nil {
		return errors.WithStack(err)
	}

	if _, err := os.Stat(filepath.Dir(absPath)); err != nil {
		err = os.MkdirAll(filepath.Dir(absPath), 0755)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if err := ioutil.WriteFile(absPath, b, 0644); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func goFmt(path string, b []byte) ([]byte, error) {
	if len(path) < 4 || path[len(path)-3:] != ".go" {
		return b, nil
	}

	fmtb, err := imports.Process(path, b, nil)
	if err != nil {
		return b, errors.WithStack(err)
	}

	return fmtb, nil
}

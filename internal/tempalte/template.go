package apptemplate

import (
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/hlts2/create-go-app/internal/tempalte/datas/app"
)

// Templater is an interface for template.
type Templater interface {
	Path() string
	Template() *template.Template
}

type templaterImpl struct {
	path string
	tpl  *template.Template
}

func (t *templaterImpl) Path() string {
	return t.path
}

func (t *templaterImpl) Template() *template.Template {
	return t.tpl
}

// ApplicationTemplates is templates of application code.
func ApplicationTemplates() []Templater {
	return []Templater{
		&templaterImpl{
			path: "README.md",
			tpl:  mustTmpl("README.md", app.ReadmeTmpl),
		},

		&templaterImpl{
			path: "go.mod",
			tpl:  mustTmpl("go.mod", app.ModTmpl),
		},

		&templaterImpl{
			path: "main.go",
			tpl:  mustTmpl("main.go", app.MainTmpl),
		},

		&templaterImpl{
			path: filepath.Join("pkg", "config", "config.go"),
			tpl:  mustTmpl("config.go", app.ConfigTmpl),
		},

		&templaterImpl{
			path: filepath.Join("pkg", "infra", "repository", "api", "repository.go"),
			tpl:  mustTmpl("api_repository.go", app.APIRepositoryTmpl),
		},

		&templaterImpl{
			path: filepath.Join("pkg", "infra", "repository", "db", "repository.go"),
			tpl:  mustTmpl("db_repository.go", app.DBRepositoryTmpl),
		},

		&templaterImpl{
			path: filepath.Join("pkg", "domain", "service", "service.go"),
			tpl:  mustTmpl("service.go", app.ServiceTmpl),
		},

		&templaterImpl{
			path: filepath.Join("pkg", "domain", "model", "model.go"),
			tpl:  mustTmpl("service.go", app.ModelTmpl),
		},

		&templaterImpl{
			path: filepath.Join("pkg", "usecase", "usecase.go"),
			tpl:  mustTmpl("usecase.go", app.UsecaseTmpl),
		},
	}
}

// CLIApplicationTemplates is templates of CLI application code.
func CLIApplicationTemplates() []Templater {
	return nil
}

var tmplFuncs = template.FuncMap{
	"ToCamel":      strcase.ToCamel,
	"ToLowerCamel": strcase.ToLowerCamel,
}

func mustTmpl(name, text string) *template.Template {
	return template.Must(template.New(name).Funcs(tmplFuncs).Parse(text))
}

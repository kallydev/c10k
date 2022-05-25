package template

import (
	"embed"
	"html/template"
)

var (
	//go:embed *.tmpl
	FS embed.FS
)

func Parse() (*template.Template, error) {
	return template.ParseFS(FS, "*.tmpl")
}

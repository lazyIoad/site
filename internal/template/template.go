package template

import (
	"io"
	tp "text/template"
)

var templates = tp.Must(tp.ParseGlob("web/templates/*.tmpl"))

func Render(w io.Writer, path string, data any) error {
	err := templates.ExecuteTemplate(w, path+".tmpl", data)
	if err != nil {
		return err
	}

	return nil
}


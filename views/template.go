package views

import (
	"fmt"
	"html/template"
	"net/http"
)

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)

	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{htmlTpl: tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	if err := t.htmlTpl.Execute(w, data); err != nil {
		fmt.Printf("failed execute template: %v\n", err)
		http.Error(w, "encounter an error executing the template", http.StatusInternalServerError)
		return
	}

}

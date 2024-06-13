package controller

import (
	"net/http"

	"github.com/devphasex/lenlocked/views"
)

func StaticHanlder(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/devphasex/lenlocked/controller"
	"github.com/devphasex/lenlocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/", controller.StaticHanlder(tpl))
	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/contact", controller.StaticHanlder(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/faq", controller.StaticHanlder(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Server running on port:3001")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handleFunc)
	r.Get("/contact", contactHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Server running on port:3001")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}

func handleFunc(w http.ResponseWriter, _ *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")

	tpl, err := template.ParseFiles(filepath)

	if err != nil {
		fmt.Printf("failed parsing template: %v\n", err)
		http.Error(w, "encounter an error parsing the template", http.StatusInternalServerError)
		return
	}

	if err = tpl.Execute(w, nil); err != nil {
		fmt.Printf("failed execute template: %v\n", err)
		http.Error(w, "encounter an error executing the template", http.StatusInternalServerError)
		return
	}
}

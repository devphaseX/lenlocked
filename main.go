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
	w.Header().Set("Content-type", "text/html; charset=utf=8")

	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)

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

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact page</h1><p>To get in touch reach out to me
	 <a href=\"mailto:ayomidelawal700@gmail.com\">ayomidelawal700@gmail.com</a></p>`)
}

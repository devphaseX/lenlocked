package main

import (
	"fmt"
	"net/http"
)

type Server struct{}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleFunc(w, r)

	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}

}

func main() {
	var server Server
	fmt.Println("Server running on port:3000")
	http.ListenAndServe(":3000", server)
}

func handleFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	fmt.Fprint(w, "<h1>Welcome to my go love</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact page</h1><p>To get in touch reach out to me
	 <a href=\"mailto:ayomidelawal700@gmail.com\">ayomidelawal700@gmail.com</a></p>`)
}

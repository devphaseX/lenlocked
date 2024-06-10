package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Server running on port:3000")
	http.ListenAndServe(":3000", nil)
}

func handleFunc(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my go love</h1>")
}

package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello, World! This is my first Awesome and Great Go web app!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}

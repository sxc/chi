package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello, World! This is my first Awesome and Great Go web app!</h1>")
	// fmt.Fprint(w, r.URL.Path)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, please send an email to <a href=\"mailto:abc@example.com\">Email</a>.</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>Sorry, but the page you were trying to view does not exist.</p>")
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	http.HandleFunc("/", pathHandler)
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", router)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>Sorry, but the page you were trying to view does not exist.</p>")
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}
}

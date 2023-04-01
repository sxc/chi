package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

// "fmt"
// "html/template"
// "log"
// "net/http"
// "path/filepath"

// "github.com/go-chi/chi/v5"
// "fmt"
// "log"
// "net/http"
// "path/filepath"
// "text/template"

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	tpl, err := template.ParseFiles("templates/home.gohtml")
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	err = tpl.Execute(w, nil)
// 	if err != nil {
// 		log.Printf("executing template: %v", err)
// 		http.Error(w, "Internal server error executing the template.", http.StatusInternalServerError)
// 	}
// }

func executeTemplate(w http.ResponseWriter, filepath string) {
	// t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	tpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		log.Printf("executing template: %v", err)
// 		http.Error(w, "Internal server error executing the template.", http.StatusInternalServerError)
// 		return
// 	}
// 	err = tpl.Execute(w, nil)
// 	if err != nil {
// 		log.Printf("executing template: %v", err)
// 		http.Error(w, "Internal server error executing the template.", http.StatusInternalServerError)
// 		return
// 	}
// }

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"

	//"os/exec"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, tplPath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "Parsing template Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, "Executing template Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
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
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

// add galleries handler
func galleriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Galleries Page</h1>")
	idHere := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("Galleries: %s", idHere)))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/galleries/{id}", galleriesHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFoundHandler().ServeHTTP(w, r)
	})
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}

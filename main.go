package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/chat.html"))
	tmpl.Execute(w, nil)
}

func mediaHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/media.html"))
	tmpl.Execute(w, nil)
}

// I love Air

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path { // returns the URL and path
	case "/":
		homeHandler(w, r)
	case "/chat":
		chatHandler(w, r)
	case "/media":
		mediaHandler(w, r)
	default:
		// Page not found :(
		http.NotFound(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", router)
}

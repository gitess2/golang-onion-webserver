package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome!</h>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", handler)
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
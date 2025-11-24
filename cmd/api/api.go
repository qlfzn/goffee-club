package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type application struct {
	addr string
	// db
	// logger
}

// create HTTP router for api
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	return r
}
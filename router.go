package main

import (
	"net/http"
	"urza/handlers"
	"urza/models"

	"github.com/gorilla/mux"
)

func buildRouter(db *models.AppDb) http.Handler {
	r := mux.NewRouter()
	env := handlers.UrzaEnvironment{DB: db}

	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// CREATION OF A SUBROUTER
	api := r.PathPrefix("/api").Subrouter()

	// EXPENSES ROUTES
	api.Handle("/expenses/{id}", handlers.Register(&env)).Methods("POST").Name("Expenses")

	return r
}

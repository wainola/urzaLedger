package main

import (
	"net/http"
	"urza/handlers"
	"urza/models"

	"github.com/gorilla/mux"
)

func buildRouter(db *models.UrzaDB) http.Handler {
	r := mux.NewRouter()
	env := handlers.UrzaEnvironment{DB: db}

	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// CREATION OF A SUBROUTER
	api := r.PathPrefix("/api").Subrouter()

	// EXPENSES ROUTES
	api.Handle("/expenses/{id}", handlers.Expenses(&env)).Methods("POST").Name("Post Expenses")
	api.Handle("/expenses/{id}", handlers.GetExpenses(&env)).Methods("GET").Name("Get Expenses")
	api.Handle("/expenses/{id}/expense/{idExpense}", handlers.EditExpense(&env)).Methods("PUT").Name("Put Expenses")

	// INCOME ROUTES
	api.Handle("/incomes/{id}", handlers.Incomes(&env)).Methods("POST").Name("Post Incomes")
	api.Handle("/incomes/{id}", handlers.GetIncomes(&env)).Methods("GET").Name("Get Incomes")

	return r
}

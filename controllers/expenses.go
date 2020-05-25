package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Expense struct {
	Date    string
	Expense string
	Amount  int
	Source  string
}

type response struct {
	Code    int
	Message string
}

func PostExpense(w http.ResponseWriter, decoder *json.Decoder) {
	fmt.Println("Post Expense")

	exp := Expense{}

	err := decoder.Decode(&exp)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := response{201, "Created"}
	json.NewEncoder(w).Encode(res)
}

func GetExpenses(w http.ResponseWriter) {
	w.WriteHeader(200)
}

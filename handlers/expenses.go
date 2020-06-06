package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"urza/utils"
)

type PostExpenseBody struct {
	Date    string
	Expense string
	Amount  int
	Source  string
}

func PostExpense(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	fmt.Println("body", body)
	bodyToValidate := PostExpenseBody{}
	err := json.NewDecoder(body).Decode(&bodyToValidate)

	if err != nil {
		panic(err)
	}

	fmt.Println("body::", bodyToValidate)
}

func ExpensesRoute(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "POST":
		utils.ValidateHeader(w, r, PostExpense)
	}

}

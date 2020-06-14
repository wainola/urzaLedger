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

type Response struct {
	Code    int
	Message string
}

type GetExpenseResponse struct {
	id        string
	expense   string
	amount    int
	source    string
	createdAt string
}

type PutExpenseBody struct {
	Id      string
	Expense string
	Amount  int
	Source  string
	Date    string
}

func Expenses(appEnvironment *UrzaEnvironment) http.Handler {
	return UrzaApp{appEnvironment, postExpense}
}

func PostExpenseValidation(p PostExpenseBody) bool {
	if p.Amount == 0 {
		return false
	} else if p.Expense == "" {
		return false
	} else if p.Source == "" {
		return false
	} else {
		return true
	}
}

func postExpense(ue *UrzaEnvironment, w http.ResponseWriter, r *http.Request) error {
	body := r.Body

	_, userId := utils.ExtractUrlToProcess(r.URL.Path, "post", "expenses")

	bodyToValidate := PostExpenseBody{}
	err := json.NewDecoder(body).Decode(&bodyToValidate)

	if err != nil {
		panic(err)
	}

	validation := PostExpenseValidation(bodyToValidate)

	if !validation {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		responseToSend := Response{http.StatusUnprocessableEntity, "Unprocessable Entity"}
		json.NewEncoder(w).Encode(responseToSend)
	}

	result := ue.DB.CreateExpense(bodyToValidate.Date, bodyToValidate.Expense, bodyToValidate.Amount, bodyToValidate.Source, userId)

	if result {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		res := Response{201, "Created"}
		json.NewEncoder(w).Encode(res)
	}

	return nil
}

func GetExpenses(appEnvironment *UrzaEnvironment) http.Handler {
	return UrzaApp{appEnvironment, getExpense}
}

func getExpense(ue *UrzaEnvironment, w http.ResponseWriter, r *http.Request) error {
	_, userId := utils.ExtractUrlToProcess(r.URL.Path, "get", "expenses")

	result := ue.DB.GetExpense(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

	return nil

}

func EditExpense(appEnvironment *UrzaEnvironment) http.Handler {
	return UrzaApp{appEnvironment, editExpense}
}

func editExpense(ue *UrzaEnvironment, w http.ResponseWriter, r *http.Request) error {
	ids, _ := utils.ExtractUrlToProcess(r.URL.Path, "put", "expenses")

	body := r.Body

	bodyToValidate := PutExpenseBody{}
	err := json.NewDecoder(body).Decode(&bodyToValidate)

	utils.HandleErr(err)

	fmt.Println("ids::", ids)

	return nil
}

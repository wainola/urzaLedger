package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"urza/utils"

	"github.com/google/uuid"
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

type ExpenseEntity struct {
	Id        string
	Expense   string
	Amount    int
	Source    string
	CreatedAt string
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

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PostExpense(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	body := r.Body

	bodyToValidate := PostExpenseBody{}
	err := json.NewDecoder(body).Decode(&bodyToValidate)

	if err != nil {
		panic(err)
	}

	validation := PostExpenseValidation(bodyToValidate)

	if !validation {
		w.Header().Set("Content-Type", "application/json")
		validation = PostExpenseValidation(bodyToValidate)
		w.WriteHeader(http.StatusUnprocessableEntity)
		responseToSend := Response{http.StatusUnprocessableEntity, "Unprocessable Entity"}
		json.NewEncoder(w).Encode(responseToSend)
	}

	defer db.Close()

	tx, err := db.Begin()

	HandleErr(err)

	expenseToInsert := ExpenseEntity{
		uuid.New().String(),
		bodyToValidate.Expense,
		bodyToValidate.Amount,
		bodyToValidate.Source,
		bodyToValidate.Date,
	}

	insertionStmt, err := tx.Prepare("insert into Expenses (id, expense, amount, source, createdAt) values (?, ?, ?, ?, ?);")

	HandleErr(err)

	defer insertionStmt.Close()

	_, err = insertionStmt.Exec(expenseToInsert.Id, expenseToInsert.Expense, expenseToInsert.Amount, expenseToInsert.Source, expenseToInsert.CreatedAt)

	HandleErr(err)

	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func ExpensesRoute(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "POST":
		utils.ValidateHeader(w, r, PostExpense)
	}

}

package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
	UserId    string
}

type GetExpenseResponse struct {
	id        string
	expense   string
	amount    int
	source    string
	createdAt string
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

	userId := strings.Split(r.URL.Path, "/")[2]

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
		userId,
	}

	insertionStmt, err := tx.Prepare("insert into Expenses (id, expense, amount, source, createdAt, user_id) values (?, ?, ?, ?, ?, ?);")

	HandleErr(err)

	defer insertionStmt.Close()

	_, err = insertionStmt.Exec(expenseToInsert.Id, expenseToInsert.Expense, expenseToInsert.Amount, expenseToInsert.Source, expenseToInsert.CreatedAt, expenseToInsert.UserId)

	HandleErr(err)

	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expenseToInsert)
}

func GetExpense(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println(r.URL)

	userId := strings.Split(r.URL.Path, "/")[2]

	defer db.Close()

	tx, err := db.Begin()

	HandleErr(err)

	getStmt, err := tx.Prepare("select id, expense, amount, source, createdAt from Expenses where user_id=?;")

	HandleErr(err)

	defer getStmt.Close()

	rows, err := getStmt.Query(userId)

	defer rows.Close()

	expenses := []GetExpenseResponse{}

	for rows.Next() {
		var id string
		var expense string
		var amount int
		var source string
		var createdAt string

		err = rows.Scan(&id, &expense, &amount, &source, &createdAt)

		HandleErr(err)

		expenses = append(expenses, GetExpenseResponse{id, expense, amount, source, createdAt})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&expenses)
	fmt.Println(expenses)

}

func ExpensesRoute(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "POST":
		utils.ValidateHeader(w, r, PostExpense)
	case "GET":
		utils.ValidateHeader(w, r, GetExpense)
	}

}

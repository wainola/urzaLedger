package models

import (
	"urza/utils"

	"github.com/google/uuid"
)

type ExpenseEntity struct {
	Id        string `json:"id"`
	Expense   string `json:"expense"`
	Amount    int    `json:"amount"`
	Source    string `json:"source"`
	CreatedAt string `json:"createdAt"`
}

func (db *UrzaDB) CreateExpense(date string, expense string, amount int, source string, userId string) bool {
	tx, err := db.DB.Begin()

	utils.HandleErr(err)

	insertionStmt, err := tx.Prepare("insert into Expenses (id, expense, amount, source, createdAt, user_id) values (?, ?, ?, ?, ?, ?);")

	utils.HandleErr(err)

	defer insertionStmt.Close()

	_, err = insertionStmt.Exec(uuid.New().String(), expense, amount, source, date, userId)

	utils.HandleErr(err)

	tx.Commit()

	return true
}

func (db *UrzaDB) GetExpense(userId string) []ExpenseEntity {
	tx, err := db.DB.Begin()

	utils.HandleErr(err)

	getStmt, err := tx.Prepare("select id, expense, amount, source, createdAt from Expenses where user_id=?;")

	utils.HandleErr(err)

	defer getStmt.Close()

	rows, err := getStmt.Query(userId)

	defer rows.Close()

	expenses := []ExpenseEntity{}

	for rows.Next() {
		var id string
		var expense string
		var amount int
		var source string
		var createdAt string

		err = rows.Scan(&id, &expense, &amount, &source, &createdAt)

		utils.HandleErr(err)

		expenses = append(expenses, ExpenseEntity{id, expense, amount, source, createdAt})
	}

	return expenses
}

func (db *UrzaDB) EditExpense(ids []string) bool {
	_, err := db.DB.Begin()

	utils.HandleErr(err)

	// putStmt, err :=  tx.Prepare("")

	return true
}

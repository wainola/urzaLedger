package models

import (
	"urza/utils"

	"github.com/google/uuid"
)

type IncomeEntity struct {
	Id        string
	Income    string
	Amount    int
	Source    string
	CreatedAt string
}

func (db *UrzaDB) CreateIncome(income string, amount int, source string, createdAt string, userId string) bool {
	tx, err := db.DB.Begin()

	utils.HandleErr(err)

	insertionStmt, err := tx.Prepare("insert into Incomes (id, income, amount, source, createdAt, user_id) values(?,?,?,?,?,?);")

	utils.HandleErr(err)

	defer insertionStmt.Close()

	_, err = insertionStmt.Exec(uuid.New().String(), income, amount, source, createdAt, userId)

	utils.HandleErr(err)

	tx.Commit()

	return true
}

func (db *UrzaDB) GetIncomes(userId string) []IncomeEntity {
	tx, err := db.DB.Begin()

	utils.HandleErr(err)

	getStmt, err := tx.Prepare("select id, income, amount, source, createdAt from Incomes where user_id=?;")

	defer getStmt.Close()

	rows, err := getStmt.Query(userId)

	defer rows.Close()

	incomes := []IncomeEntity{}

	for rows.Next() {
		var id string
		var income string
		var amount int
		var source string
		var createdAt string

		err = rows.Scan(&id, &income, &amount, &source, &createdAt)

		utils.HandleErr(err)

		incomes = append(incomes, IncomeEntity{id, income, amount, source, createdAt})
	}

	return incomes
}

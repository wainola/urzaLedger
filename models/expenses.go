package models

type ExpenseEntity struct {
	Id        string
	Expense   string
	Amount    int
	Source    string
	CreatedAt string
	UserId    string
}

func (db *UrzaDB) CreateExpense() 

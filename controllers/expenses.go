package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Expense struct {
	Date    string
	Expense string
	Amount  int
	Source  string
}

type InMemoriExpenses struct {
	store []Expense
	lock  sync.RWMutex
}

func NewInMemoryExpense() *InMemoriExpenses {
	return &InMemoriExpenses{
		[]Expense{},
		sync.RWMutex{},
	}
}

type response struct {
	Code    int
	Message string
}

func PostExpense(w http.ResponseWriter, decoder *json.Decoder, expensesStore *InMemoriExpenses) {
	fmt.Println("Post Expense")

	exp := Expense{}

	err := decoder.Decode(&exp)

	if err != nil {
		panic(err)
	}

	expensesStore.lock.Lock()
	defer expensesStore.lock.Unlock()
	expensesStore.store = append(expensesStore.store, exp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := response{201, "Created"}
	json.NewEncoder(w).Encode(res)
}

func GetExpenses(w http.ResponseWriter, expensesStore *InMemoriExpenses) {
	fmt.Println(expensesStore.store)
	w.WriteHeader(200)
}

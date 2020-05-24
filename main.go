package main

import (
	"fmt"
	"net/http"

	"./controllers"
	"./handlers"

	"github.com/gorilla/mux"
)

func main() {
	expensesStore := controllers.NewInMemoryExpense()

	r := mux.NewRouter()
	r.HandleFunc("/expense", handlers.HandleExpense)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}

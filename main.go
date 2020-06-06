package main

import (
	"fmt"
	"net/http"
	"urza/handlers"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/expense", handlers.ExpensesRoute)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}

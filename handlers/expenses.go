package handlers

import (
	"fmt"
	"net/http"

	"../controllers"
)

func HandleExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleExpense")

	method := r.Method

	switch method {
	case "POST":
		controllers.PostExpense(w, r)
	default:
		w.WriteHeader(503)
		w.Write([]byte("Service unavailable"))
	}
}

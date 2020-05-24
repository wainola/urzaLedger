package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../controllers"
)

func HandleExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleExpense")

	method := r.Method
	// decode json body
	decoder := json.NewDecoder(r.Body)

	switch method {
	case "POST":
		controllers.PostExpense(w, decoder)
	default:
		w.WriteHeader(503)
		w.Write([]byte("Service unavailable"))
	}
}

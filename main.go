package main

import (
	"fmt"
	"net/http"
	"strings"

	"./handlers"

	"github.com/gorilla/mux"
)

// func goDotEnvVariable(key string) string {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading the env file")
// 	}

// 	return os.Getenv(key)
// }

func joinStrings(strs ...string) string {
	var finalString strings.Builder

	for _, str := range strs {
		finalString.WriteString(str)
	}

	return finalString.String()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/expense", handlers.HandleExpense)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}

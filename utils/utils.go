package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Code    int
	Message string
}

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading the env file")
	}

	return os.Getenv(key)
}

func ValidateHeader(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request)) {
	header := r.Header.Get("client-name")

	if header != "app-mobile" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		forbidenRes := Response{
			403,
			"Forbidden",
		}
		json.NewEncoder(w).Encode(forbidenRes)
	}

	f(w, r)
}

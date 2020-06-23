package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

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

func ResolveResponse(res Response, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(res)
}

func ValidateHeader(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request, db *sql.DB)) {
	header := r.Header.Get("client-name")

	if header != "app-mobile" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		forbidenRes := Response{
			403,
			"Forbidden",
		}
		ResolveResponse(forbidenRes, w)
	}

	dbInstance := InstanceDbConnection()

	f(w, r, dbInstance)
}

func InstanceDbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./urza.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserId(routePrefix string, url string) string {
	s := strings.Split(url, "/")

	index := getIndex(s, routePrefix)

	return s[index]
}

func getIndex(splitedUrl []string, routePrefix string) int {
	var indextToSearch int
	for idx, v := range splitedUrl {
		if v == routePrefix {
			indextToSearch = idx
			break
		}
	}

	return indextToSearch + 1
}

// MAP OF THE CURRENT URLS IN USE FOR THE API
var urlsMap = map[string][]string{
	"put": []string{
		"api/expenses/:id/expense/:idExpense",
		"api/incomes/:id/income/:idIncome",
	},
	"post": []string{
		"api/expenses/:id",
		"api/incomes/:id",
	},
	"get": []string{
		"api/expenses/:id",
		"api/incomes/:id",
	},
}

// FUNCTION THAT RETURNS AN ARRAY OF THE IDS OF THE URLS OR JUST THE ID
// IF THE URL HAS ONLY ONE ID PARAMETER
func ExtractUrlToProcess(url string, method string, endpointDescriptor string) ([]string, string) {
	urls := urlsMap[method]
	var idx int

	for i, url := range urls {
		if strings.Contains(url, endpointDescriptor) {
			idx = i
			break
		}
	}

	urlToProcess := urls[idx]
	idsPositioning := []int{}

	for idx, str := range strings.Split(urlToProcess, "/") {
		if strings.Contains(str, ":") {
			idsPositioning = append(idsPositioning, idx)
		}
	}

	return getUserId(url, idsPositioning)
}

// FUNCTION THAT RETURNS THE IDS THAT ARE INSIDE A URL
func getUserId(url string, positions []int) ([]string, string) {
	s := strings.Split(url, "/")

	s2 := removeEmptyStrings(s)

	urlId := []string{}

	for _, position := range positions {
		urlId = append(urlId, s2[position])
	}
	if len(urlId) != 1 {
		return urlId, ""
	}
	return []string{}, urlId[0]
}

// CLEAN THE SLICES GENERATED BY THE SPLIT METHOD
func removeEmptyStrings(str []string) []string {
	newStrSlice := []string{}

	for _, v := range str {
		if v != "" {
			newStrSlice = append(newStrSlice, v)
		}
	}

	return newStrSlice
}

// FUNCTION THAT VALIDATE THE STRUCTURE OF THE BODY
func ValidateBodyStruct(value reflect.Value) bool {
	validationResult := []bool{}
	for i := 0; i < value.NumField(); i++ {
		fieldType := value.Field(i).Type()
		fieldValue := value.Field(i).Interface()
		res := checkAgainstDefaultValues(fieldType, fieldValue, value.Field(i))
		fmt.Println("field name", value.Field(i))
		validationResult = append(validationResult, res)
	}

	var lazyFlag bool

	for _, value := range validationResult {
		if !value {
			lazyFlag = value
			break
		}
	}
	return lazyFlag
}

// CONSTANT THAT DEFINE THE DEFAULT VALUES FOR SOME PRIMITIVES
// THAT ARE IN USE FOR THE API
const (
	defaultString  string = ""
	defaultInt     int    = 0
	defaultBoolean bool   = false
)

func checkAgainstDefaultValues(fType reflect.Type, fValue interface{}, field reflect.Value) bool {
	switch fType.Kind() {
	case reflect.String:
		return fValue != defaultString
	case reflect.Int:
		return fValue != defaultInt
	case reflect.Bool:
		return fValue != defaultBoolean
	case reflect.Slice:
		return field.Len() != 0
	}

	return false
}

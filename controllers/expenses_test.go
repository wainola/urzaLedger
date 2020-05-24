package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostExpense(t *testing.T) {

	data := Expense{
		"2020-05-23",
		"some payment",
		3000,
		"TP",
	}
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}
	transformedData := strings.NewReader(bytes)

	fmt.Println("transformedData", transformedData)

	req := httptest.NewRequest("POST", "/expense", transformedData)

	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()

	PostExpense(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

}

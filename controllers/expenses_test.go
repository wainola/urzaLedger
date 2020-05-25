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

	data := strings.NewReader(`{"date": "2020-05-23", "expense": "pago de repuestos", "amount": 150000, "source": "tarjeta platinum" }`)

	decoder := json.NewDecoder(data)

	w := httptest.NewRecorder()

	PostExpense(w, decoder)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(strings.Contains(string(body), "Created"))

	if resp.StatusCode != 201 {
		t.Errorf("PostExpense controler: got %d want 201", resp.StatusCode)
	}

	if !strings.Contains(string(body), "Created") {
		t.Errorf(`PostExpense response body: got %s want "{"Code":201,"Message":"Created"}"`, body)
	}

}

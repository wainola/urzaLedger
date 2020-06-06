package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleExpense(t *testing.T) {
	buf := new(bytes.Buffer)

	data := `{"date": "2020-05-23", "expense": "pago de repuestos", "amount": 150000 }`

	buf.Write([]byte(data))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/expense", buf)

	HandleExpense(w, req)
}

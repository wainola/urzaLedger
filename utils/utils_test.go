package utils

import (
	"testing"
)

func TestExtractUrlToProcess(t *testing.T) {
	t.Run("Should return the id for a url with one id as a parameter", func(t *testing.T) {
		urlToTest := "api/incomes/ABC123"
		method := "post"
		endpointDescriptor := "incomes"

		_, r2 := ExtractUrlToProcess(urlToTest, method, endpointDescriptor)

		if r2 != "ABC123" {
			t.Errorf("Error: expected %s => have %s", "ABC123", r2)
		}
	})

	t.Run("Should return the ids of the url", func(t *testing.T) {
		url := "api/incomes/ABC123/income/123XYZ"
		method := "put"
		endpoint := "incomes"

		r1, _ := ExtractUrlToProcess(url, method, endpoint)

		first := r1[0]
		second := r1[1]

		if first != "ABC123" {
			t.Errorf("Error: expected ABC123 => have %s", first)
		}

		if second != "123XYZ" {
			t.Errorf("Error: expected 123XYZ => have %s", second)
		}
	})
}

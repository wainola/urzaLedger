package utils

import (
	"fmt"
	"testing"
)

func TestExtractUrlToProcess(t *testing.T) {
	urlToTest := "some/url/to/test"
	method := "POST"
	endpointDescriptor := "expenses"

	r1, r2 := ExtractUrlToProcess(urlToTest, method, endpointDescriptor)

	fmt.Println(r1, r2)
}

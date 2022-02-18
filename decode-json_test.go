package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"firstname":"Billy", "Age":23}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}

	fmt.Println(customer.LastName)
}

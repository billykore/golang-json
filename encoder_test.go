package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	writer, _ := os.Create("customer.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Evanbill",
		MiddleName: "Antonio",
		LastName:   "Kore",
		Age:        23,
		Married:    false,
		Hobbies:    nil,
		Addresses:  nil,
	}

	err := encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
}

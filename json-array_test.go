package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Evanbill",
		MiddleName: "Antonio",
		LastName:   "Kore",
		Age:        23,
		Married:    false,
		Hobbies:    []string{"Makan", "Tidur", "Makan lagi"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonArrayUnmarshal(t *testing.T) {
	jsonString := `{"FirstName":"Evanbill","MiddleName":"Antonio","LastName":"Kore","Age":23,"Married":false,"Hobbies":["Makan","Tidur","Makan lagi"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}

	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Age)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Billy",
		MiddleName: "",
		LastName:   "",
		Age:        23,
		Married:    false,
		Hobbies:    nil,
		Addresses: []Address{
			{
				Street:     "Jalan 1",
				Country:    "Indonesia",
				PostalCode: "11111",
			},
			{
				Street:     "Jalan 2",
				Country:    "Indonesia",
				PostalCode: "22222",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Billy","MiddleName":"","LastName":"","Age":23,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan 1","Country":"Indonesia","PostalCode":"11111"},{"Street":"Jalan 2","Country":"Indonesia","PostalCode":"22222"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}

	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Addresses)
}

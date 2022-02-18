package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id": 1, "name": "MacBook", "price": "20000000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    1,
		"name":  "Macbook",
		"price": "20000000",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

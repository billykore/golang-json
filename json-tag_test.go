package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "MacBook Pro",
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

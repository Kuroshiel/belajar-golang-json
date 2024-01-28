package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Map Golang JSON

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001", "name":"Lenovo Thinkpad T410", "price":"2000000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	products := map[string]interface{}{
		"id":    "P0001",
		"name":  "Lenovo Thinkpad T410",
		"price": 2000000,
	}
	bytes, _ := json.Marshal(products)
	fmt.Println(string(bytes))
}

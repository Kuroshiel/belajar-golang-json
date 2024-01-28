package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Tag Golang JSON

type Products struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	products := Products{
		Id:       "P0001",
		Name:     "Lenovo Thinkpad T410",
		ImageUrl: "http://example.com/image.png",
	}
	bytes, _ := json.Marshal(products)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Lenovo Thinkpad T410","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	products := &Products{}

	err := json.Unmarshal(jsonBytes, products)
	if err != nil {
		panic(err)
	}

	fmt.Println(products)
	fmt.Println(products.Id)
	fmt.Println(products.Name)
	fmt.Println(products.ImageUrl)
}

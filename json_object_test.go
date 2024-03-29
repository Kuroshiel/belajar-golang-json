package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Array Complex Golang JSON

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

// JSON Object Golang Web

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        30,
		Married:    true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

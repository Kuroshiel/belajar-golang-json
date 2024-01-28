package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Array Golang JSON

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Hobbies:    []string{"Game", "Reading", "Coding"},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":0,"Married":false,"Hobbies":["Game","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

// JSON Array Complex Golang JSON

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Eko",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indoneisa",
				PostalCode: "9999",
			},

			{
				Street:     "Jalan Belum Dibangun",
				Country:    "Indoneisa",
				PostalCode: "8888",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indoneisa","PostalCode":"9999"},{"Street":"Jalan Belum Dibangun","Country":"Indoneisa","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

// JSON Array Complex Tipe data Slice Golang JSON

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indoneisa","PostalCode":"9999"},{"Street":"Jalan Belum Dibangun","Country":"Indoneisa","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indoneisa",
			PostalCode: "9999",
		},

		{
			Street:     "Jalan Belum Dibangun",
			Country:    "Indoneisa",
			PostalCode: "8888",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}

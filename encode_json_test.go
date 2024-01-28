package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Encode JSON Golang JSON

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncdo(t *testing.T) {
	logJson("Eko")
	logJson(1)
	logJson(true)
	logJson([]string{"Eko", "Kurniawan", "Khannedy"})
}

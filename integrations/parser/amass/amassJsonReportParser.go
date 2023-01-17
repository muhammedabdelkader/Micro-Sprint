package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type AmassRecord struct {
	Domain string `json:"domain"`
	IP     string `json:"ip"`
	Source string `json:"source"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("amass_report.json")
	if err != nil {
		fmt.Println(err)
		return

	}
	defer file.Close()

	// Decode the JSON into a slice of AmassRecord structs
	var records []AmassRecord
	if err := json.NewDecoder(file).Decode(&records); err != nil {
		fmt.Println(err)
		return

	}

	// Iterate through the records and print them
	for _, record := range records {
		fmt.Println(record)

	}

}

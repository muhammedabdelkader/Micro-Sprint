package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("amass_report.csv")
	if err != nil {
		fmt.Println(err)
		return

	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read in all of the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return

	}

	// Iterate through the records and print
	// them
	for _, record := range records {
		fmt.Println(record)

	}

}

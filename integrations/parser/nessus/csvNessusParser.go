package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return

	}

	// Print the header row
	fmt.Println(records[0])

	// Print the data rows
	for i, record := range records[1:] {
		fmt.Println("Record", i+1)
		for j, field := range record {
			fmt.Printf("  Field %d: %s\n", j+1, field)

		}

	}

}

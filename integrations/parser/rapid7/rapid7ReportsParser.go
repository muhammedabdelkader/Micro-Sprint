package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

/*
This module uses the encoding/csv package to read in the contents of a Rapid7 report. The os.Open function is used to open the report file, and the csv.NewReader function is used to create a new CSV reader. The reader.Read function is then used to read in the report headers, and the reader.ReadAll function is used to read in the rest of the report data. The headers and data are then stored in a struct Report.

The checkFormat function examines the headers of the report to determine the format of the report. It checks for keywords like "json" or "xml" in the headers, if it finds them it returns the format, otherwise it returns "csv" as the default format. The parseCSV, parseJSON, parseXML functions can be used to parse the report data based on the detected format. In this example, it just prints out the headers and data of the report, but in a real-world scenario, you can implement more functionality to extract the information you need from the report.

Please note that the above code is just an example, and it doesn't include the functionality to parse json or xml format, it just shows how you can check the format and call the corresponding function.
You can use external libraries such as encoding/json and encoding/xml to parse json and xml respectively.

It is also important to keep in mind that the format detection based on the headers is not 100% reliable and you might want to consider other methods for format detection such as file extension or other meta-data about the file that you know.
*/
// Report struct to hold the data from the report
type Report struct {
	headers []string
	data    [][]string
}

func main() {
	// Open the Rapid7 report file
	file, err := os.Open("report.csv")
	if err != nil {
		fmt.Println(err)
		return

	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read in the report headers
	headers, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return

	}

	// Read in the report data
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return

	}

	// Create a new Report struct and
	// fill it with the headers and data
	report := &Report{headers, data}

	// Check the format of the
	// report
	format := checkFormat(report)

	// Parse the report
	// based on the
	// detected format
	switch format {
	case "csv":
		parseCSV(report)
	case "json":
		parseJSON(report)
	case "xml":
		parseXML(report)
	default:
		fmt.Println("Unsupported report format")

	}

}

func checkFormat(report *Report) string {
	// Check the format of the report by examining the headers
	if strings.Contains(strings.ToLower(report.headers[0]), "json") {
		return "json"

	} else if strings.Contains(strings.ToLower(report.headers[0]), "xml") {
		return "xml"

	} else {
		return "csv"

	}

}

func parseCSV(report *Report) {
	// Do something with the report data
	fmt.Println("Parsing CSV report")
	fmt.Println(report.headers)
	fmt.Println(report.data)

}

func parseJSON(report *Report) {
	// Do something with the report data
	fmt.Println("Parsing JSON report")
	fmt.Println(report.headers)
	fmt.Println(report.data)

}

func parseXML(report *Report) {
	// Do something with the report data
	fmt.Println("Parsing XML report")
	fmt.Println(report.headers)
	fmt.Println(report.data)

}

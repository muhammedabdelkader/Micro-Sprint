package main

import (
	"fmt"
	"os"

	"github.com/tenable/nessus-go"
)

func parseNessus(file *os.File, format string) {
	var scan *nessus.Scan
	var err error

	switch format {
	case "nessus":
		scan, err = nessus.NewScan(file)
	case "nbe":
		scan, err = nessus.NewNBE(file)
	case "csv":
		scan, err = nessus.NewCSV(file)
	default:
		fmt.Println("Invalid format")
		return

	}

	if err != nil {
		fmt.Println("Error creating scan:", err)
		return

	}

	fmt.Println("Report name:", scan.Info.Name)
	fmt.Println("Number of hosts:", len(scan.Hosts))
	for _, host := range scan.Hosts {
		fmt.Println("Host IP:", host.Name)
		for _, item := range host.Items {
			fmt.Println("  Item:", item.Name)

		}

	}

}

func main() {
	file, err := os.Open("example.nessus")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer file.Close()

	parseNessus(file, "nessus")

	file, err = os.Open("example.nbe")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer file.Close()

	parseNessus(file, "nbe")

	file, err = os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer file.Close()

	parseNessus(file, "csv")

}

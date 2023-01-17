package main

import (
	"fmt"
	"os"

	"github.com/tenable/nessus-go"
)

func main() {
	file, err := os.Open("example.nessus")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer file.Close()

	scan, err := nessus.NewScan(file)
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

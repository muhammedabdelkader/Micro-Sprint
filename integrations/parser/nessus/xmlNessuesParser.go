package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Report struct {
	Hosts []Host `xml:"ReportHost"`
}

type Host struct {
	Name  string `xml:"name"`
	Items []Item `xml:"ReportItem"`
}

type Item struct {
	Name string `xml:"pluginName"`
}

func main() {
	file, err := os.Open("example.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return

	}

	var report Report
	err = xml.Unmarshal(data, &report)
	if err != nil {
		fmt.Println("Error unmarshalling xml:", err)
		return

	}

	fmt.Println("Number of hosts:", len(report.Hosts))
	for _, host := range report.Hosts {
		fmt.Println("Host IP:", host.Name)
		for _, item := range host.Items {
			fmt.Println("  Item:", item.Name)

		}

	}

}

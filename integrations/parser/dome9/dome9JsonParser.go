package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
This example defines a Report struct that has a slice of Finding structs, which each have fields for the region, resource, issue, recommendation and risk. The main function reads in the JSON file, unmarshals it into a Report struct, and then iterates through the findings to print out the values of each field.
Please note that this code is just an example, it may not work as is, and you may need to adjust the struct fields and the json file name accordingly.
*/

type Report struct {
	Findings []Finding `json:"findings"`
}

type Finding struct {
	Region         string `json:"region"`
	Resource       string `json:"resource"`
	Issue          string `json:"issue"`
	Recommendation string `json:"recommendation"`
	Risk           string `json:"risk"`
}

func main() {
	jsonFile, err := ioutil.ReadFile("dome9-report.json")
	if err != nil {
		fmt.Println(err)

	}

	var report Report
	json.Unmarshal(jsonFile, &report)

	for _, finding := range report.Findings {
		fmt.Println("Region:", finding.Region)
		fmt.Println("Resource:", finding.Resource)
		fmt.Println("Issue:", finding.Issue)
		fmt.Println("Recommendation:", finding.Recommendation)
		fmt.Println("Risk:", finding.Risk)
		fmt.Println("-------------------")

	}

}

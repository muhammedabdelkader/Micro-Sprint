package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
This example defines a Report struct that has fields for the Target, ScannedAt, and an array of Vulnerability structs, which each have fields
*/
type Report struct {
	Target          string          `json:"Target"`
	ScannedAt       string          `json:"ScannedAt"`
	Vulnerabilities []Vulnerability `json:"Vulnerabilities"`
}

type Vulnerability struct {
	VulnerabilityID  string   `json:"VulnerabilityID"`
	PkgName          string   `json:"PkgName"`
	InstalledVersion string   `json:"InstalledVersion"`
	FixedVersion     string   `json:"FixedVersion"`
	Title            string   `json:"Title"`
	Severity         string   `json:"Severity"`
	Description      string   `json:"Description"`
	References       []string `json:"References"`
	CVSS             string   `json:"CVSS"`
}

func main() {
	jsonFile, err := ioutil.ReadFile("trivy-report.json")
	if err != nil {
		fmt.Println(err)

	}

	var report Report
	json.Unmarshal(jsonFile, &report)

	fmt.Println("Target:", report.Target)
	fmt.Println("Scanned At:", report.ScannedAt)

	for _, vuln := range report.Vulnerabilities {
		fmt.Println("VulnerabilityID:", vuln.VulnerabilityID)
		fmt.Println("PkgName:", vuln.PkgName)
		fmt.Println("InstalledVersion:", vuln.InstalledVersion)
		fmt.Println("FixedVersion:", vuln.FixedVersion)
		fmt.Println("Title:", vuln.Title)
		fmt.Println("Severity:", vuln.Severity)
		fmt.Println("Description:", vuln.Description)
		fmt.Println("References:", vuln.References)
		fmt.Println("CVSS:", vuln.CVSS)
		fmt.Println("-------------------")

	}

}

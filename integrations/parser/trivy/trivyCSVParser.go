package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
This code creates a new CSV file and writes the header and data of each vulnerability, after that, it iterates through the vulnerabilities and prints out the values of each field, just as it was doing before.
It is important to note that in the case of CSV parsing, the References field is a slice of strings, and in this example I'm assuming that the slice have just one element, and accessing the first element of the slice. Therefore, you may need to adjust the code to handle the case where the slices have more than one element.
Please note that this code is just an example, it may not work as is, and you may need to adjust the struct fields and the file names accordingly.
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

	// CSV parsing
	csvFile, err := os.Create("trivy-report.csv")
	if err != nil {
		fmt.Println(err)

	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	csvWriter.Write([]string{"VulnerabilityID", "PkgName", "InstalledVersion", "FixedVersion", "Title", "Severity", "Description", "References", "CVSS"})
	for _, vuln := range report.Vulnerabilities {
		csvWriter.Write([]string{vuln.VulnerabilityID, vuln.PkgName, vuln.InstalledVersion, vuln.FixedVersion, vuln.Title, vuln.Severity, vuln.Description, vuln.References[0], vuln.CVSS})

	}

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

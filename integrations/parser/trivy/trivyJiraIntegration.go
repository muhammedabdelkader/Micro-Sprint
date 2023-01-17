package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
This code iterates through the vulnerabilities in the Trivy report and for each vulnerability, it creates a new issue in Jira using the Jira REST API, it sets the project, summary, description, issuetype and priority fields of the issue with the corresponding values from the vulnerability struct.
You will need to replace YOUR_PROJECT_KEY, YOUR_JIRA_URL, and YOUR_AUTHORIZATION with the appropriate values for your Jira instance.
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

	for _, vuln := range report.Vulnerabilities {
		// Create the issue in Jira
		jsonValue, _ := json.Marshal(map[string]string{
			"fields": map[string]interface{}{
				"project": map[string]string{
					"key": "YOUR_PROJECT_KEY",
				},
				"summary":     vuln.Title,
				"description": vuln.Description,
				"issuetype": map[string]string{
					"name": "Bug",
				},
				"priority": map[string]string{
					"name": vuln.Severity,
				},
			},
		})

		client := &http.Client{}
		req, _ := http.NewRequest("POST", "https://YOUR_JIRA_URL/rest/api/2/issue", bytes.NewBuffer(jsonValue))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Basic YOUR_AUTHORIZATION")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)

		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

	}

}

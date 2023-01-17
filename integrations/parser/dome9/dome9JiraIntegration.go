package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/andygrunwald/go-jira"
)

/*
This example uses the go-jira library to interact with Jira API. The main function reads in the JSON file, unmarshals it into a Report struct, and then iterates through the findings. For each finding, it creates a new Jira issue with the details of the finding, such as region, resource, issue, recommendation and risk.
Make sure you have to replace the url, username and password with your jira instance details and adjust the project key, issue type, and priority accordingly.
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

	// Jira integration
	jiraClient, err := jira.NewClient(nil, "https://your-jira-instance.com")
	if err != nil {
		fmt.Println(err)

	}
	jiraClient.Authentication.SetBasicAuth("username", "password")

	for _, finding := range report.Findings {
		issue := jira.Issue{
			Fields: &jira.IssueFields{
				Project: jira.Project{
					Key: "SEC",
				},
				Summary:     fmt.Sprintf("%s - %s", finding.Region, finding.Issue),
				Description: fmt.Sprintf("Resource: %s\nRecommendation: %s\nRisk: %s", finding.Resource, finding.Recommendation, finding.Risk),
				Type: jira.IssueType{
					Name: "Task",
				},
				Priority: jira.Priority{
					Name: "High",
				},
			},
		}
		_, _, err := jiraClient.Issue.Create(&issue)
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println("Successfully created issue for finding: ", finding.Issue)

	}

}

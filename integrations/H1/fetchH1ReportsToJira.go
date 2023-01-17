package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andygrunwald/go-jira"
)

/*
This code example shows how to make a GET request to the HackerOne API to retrieve a report with a specific ID, and then create a new Jira issue with the title and description of the report.
You'll need to replace YOUR_API_KEY, YOUR_USERNAME, and YOUR_PASSWORD with your actual API key, username, and password.
*/
func main() {
	// Set up the HackerOne API client
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.hackerone.com/v1/reports/123", nil)
	req.Header.Add("Authorization", "Bearer YOUR_API_KEY")

	// Make the request to HackerOne
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Unmarshal the JSON response
	var report map[string]interface{}
	json.Unmarshal(body, &report)

	// Set up the Jira client
	jiraClient, _ := jira.NewClient(nil, "https://your-jira-instance.com")
	jiraClient.Authentication.SetBasicAuth("YOUR_USERNAME", "YOUR_PASSWORD")

	// Create a new
	// Jira issue
	issue := &jira.Issue{
		Fields: &jira.IssueFields{
			Summary:     report["title"].(string),
			Description: report["description"].(string),
			Type: jira.IssueType{
				Name: "Bug",
			},
		},
	}

	_, _, err := jiraClient.Issue.Create(issue)
	if err != nil {
		fmt.Println(err)

	}

}

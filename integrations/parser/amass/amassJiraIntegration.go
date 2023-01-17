package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

/*
This example demonstrates how to create a new Jira client, set the basic auth credentials, and get the current user, as well as how to create a new issue. Note that you should replace "https://yourjira.com" with the URL of your Jira instance and "username" and "password" with your Jira username and password.

You can also search for issues, update issues, add comments, and perform other operations using the methods available in the Jira client. You can find the full documentation for the "github.com/andygrunwald/go-jira" library here: https://godoc.org/github.com/andygrunwald/go-jira#Client
*/
func main() {
	// Create a new Jira client
	jiraClient, err := jira.NewClient(nil, "https://yourjira.com")
	if err != nil {
		fmt.Println(err)
		return

	}

	// Set the basic auth credentials
	jiraClient.Authentication.SetBasicAuth("username", "password")

	// Get the current user
	user, _, err := jiraClient.User.GetSelf()
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Printf("Logged in as: %s\n", user.Name)

	// Create a new issue
	issue := &jira.Issue{
		Fields: &jira.IssueFields{
			Description: "This is a test issue",
			Project:     jira.Project{Key: "TEST"},
			Summary:     "Test issue",
			Type:        jira.IssueType{Name: "Bug"},
		},
	}
	newIssue, _, err := jiraClient.Issue.Create(issue)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Printf("Successfully created issue: %s\n", newIssue.Key)

}

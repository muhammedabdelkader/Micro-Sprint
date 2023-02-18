package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Jira struct holds the information for the Jira connection
type Jira struct {
	Username string
	Password string
	BaseUrl  string
}

// Issue struct holds the information for the Jira issue
type Issue struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		Status  struct {
			Name string `json:"name"`
		} `json:"status"`
		Assignee struct {
			Name string `json:"name"`
		} `json:"assignee"`
	} `json:"fields"`
}

// Project struct holds the information for the Jira project
type Project struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func (j *Jira) SetAuth(username, password string) {
	j.Username = username
	j.Password = password

}

func (j *Jira) SetBaseUrl(baseUrl string) {
	j.BaseUrl = baseUrl

}

// GetIssue returns the information for a specific issue
func (j *Jira) GetIssue(key string) (*Issue, error) {
	req, err := http.NewRequest("GET", j.BaseUrl+"/rest/api/2/issue/"+key, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(j.Username, j.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Jira returned status %d: %s", resp.StatusCode, string(body))
	}
	var issue Issue
	if err := json.Unmarshal(body, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// UpdateIssue updates the fields of a specific issue
func (j *Jira) UpdateIssue(key string, fields map[string]interface{}) error {
	jsonData, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", j.BaseUrl+"/rest/api/2/issue/"+key, json.NewReader(jsonData))
	if err != nil {
		return err
	}
	req.SetBasicAuth(j.Username, j.Password)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Jira returned status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

// ListIssuesInProject returns a list of issues for a specific project
func (j *Jira) ListIssuesInProject(projectKey string) ([]*Issue, error) {
	req, err := http.NewRequest("GET", j.BaseUrl+"/rest/api/2/search?jql=project="+projectKey, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(j.Username, j.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Jira returned status %d: %s", resp.StatusCode, string(body))
	}
	var result struct {
		Issues []*Issue `json:"issues"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Issues, nil
}

// ListProjects returns a list of projects
func (j *Jira) ListProjects() ([]*Project, error) {
	req, err := http.NewRequest("GET", j.BaseUrl+"/rest/api/2/project", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(j.Username, j.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Jira returned status %d: %s", resp.StatusCode, string(body))
	}
	var result []*Project
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// AddProject creates a new project
func (j *Jira) AddProject(key, name string) (*Project, error) {
	data := map[string]interface{}{
		"key":  key,
		"name": name,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", j.BaseUrl+"/rest/api/2/project", json.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(j.Username, j.Password)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Jira returned status %d: %s", resp.StatusCode, string(body))
	}
	var project Project
	if err := json.Unmarshal(body, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// AddIssueInProject creates a new issue in a specific project
func (j *Jira) AddIssueInProject(projectKey, summary string) (*Issue, error) {
	data := map[string]interface{}{
		"fields": map[string]interface{}{
			"project": map[string]interface{}{
				"key": projectKey,
			},
			"summary": summary,
			"issuetype": map[string]interface{}{
				"name": "Bug",
			},
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", j.BaseUrl+"/rest/api/2/issue", json.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(j.Username, j.Password)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Jira returned status %d: %s", resp.StatusCode, string(body))
	}
	var issue Issue
	if err := json.Unmarshal(body, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// AssignIssue assigns a specific issue to a user
func (j *Jira) AssignIssue(key, username string) error {
	data := map[string]interface{}{
		"name": username,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", j.BaseUrl+"/rest/api/2/issue/"+key+"/assignee", json.NewReader(jsonData))
	if err != nil {
		return err
	}
	req.SetBasicAuth(j.Username, j.Password)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Jira returned status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

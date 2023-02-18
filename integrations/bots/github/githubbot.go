package github

import (
	"context"

	"github.com/google/go-github/github"
)

/*
* The github package exports a Client struct which has a GithubClient field of
* type *github.Client. The NewClient function returns a new Client struct with a
* properly instantiated GithubClient field. The ListRepos, CreateIssue,
* SubmitComment, and ReplyComment methods are provided as an example of how to
* use the GithubClient to interact with the GitHub API to list repos, create
* issues, submit comments, and reply to comments
 */
// Client struct to hold the github client
type Client struct {
	GithubClient *github.Client
}

// NewClient returns a new github client
func NewClient(token string) *Client {
	var client *github.Client
	if token != "" {
		ts := github.BasicAuthTransport{
			Username: "",
			Password: token,
		}
		client = github.NewClient(ts.Client())

	} else {
		client = github.NewClient(nil)

	}
	return &Client{client}

}

// ListRepos lists the repos of a user
func (c *Client) ListRepos(username string) ([]*github.Repository, error) {
	repos, _, err := c.GithubClient.Repositories.List(context.Background(), username, nil)
	if err != nil {
		return nil, err

	}
	return repos, nil

}

// CreateIssue creates an issue in a repo
func (c *Client) CreateIssue(owner, repo string, issue *github.IssueRequest) (*github.Issue, error) {
	issue, _, err := c.GithubClient.Issues.Create(context.Background(), owner, repo, issue)
	if err != nil {
		return nil, err

	}
	return issue, nil

}

// SubmitComment submit a comment to an issue
func (c *Client) SubmitComment(owner, repo string, issueNumber int, comment string) (*github.IssueComment, error) {
	commentReq := &github.IssueComment{
		Body: &comment,
	}
	comment, _, err := c.GithubClient.Issues.CreateComment(context.Background(), owner, repo, issueNumber, commentReq)
	if err != nil {
		return nil, err

	}
	return comment, nil

}

// ReplyComment reply to a comment on an issue
func (c *Client) ReplyComment(owner, repo string, commentID int, comment string) (*github.IssueComment, error) {
	commentReq := &github.IssueComment{
		Body: &comment,
	}
	comment, _, err := c.GithubClient.Issues.EditComment(context.Background(), owner, repo, commentID, commentReq)
	if err != nil {
		return nil, err

	}
	return comment, nil

}

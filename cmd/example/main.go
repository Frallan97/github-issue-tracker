package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Frallan97/github-issue-tracker/pkg/issue"
)

func main() {
	// Get GitHub PAT from environment
	pat := os.Getenv("GITHUB_PAT")
	if pat == "" {
		log.Fatal("GITHUB_PAT environment variable is required")
	}

	// Create new issue service
	service := issue.New(issue.Config{
		PATToken:   pat,
		Owner:      "your-username",
		Repo:       "your-repo",
		HTTPClient: http.DefaultClient,
	})

	// Create a new issue
	newIssue := &issue.Issue{
		Title: "Test Issue",
		Body:  "This is a test issue created via the API",
		Labels: []string{
			"test",
			"api",
		},
	}

	response, err := service.Create(&issue.IssueRequest{Issue: newIssue})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created issue #%d: %s\n", response.Number, response.HTMLURL)

	// Get issue by number
	fetchedIssue, err := service.Get(response.Number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved issue: %s\n", fetchedIssue.Title)

	// Update the issue
	updateIssue := &issue.Issue{
		Title: "Updated Test Issue",
		State: "closed",
	}

	updated, err := service.Update(response.Number, &issue.IssueRequest{Issue: updateIssue})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated issue: %s (Status: %s)\n", updated.Title, updated.State)
}

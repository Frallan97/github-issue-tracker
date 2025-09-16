package main

import (
	"fmt"
	"log"
	"os"

	"github.com/frans-sjostrom/github-issue-tracker/pkg/issue"
)

func main() {
	// Get GitHub configuration from environment
	pat := os.Getenv("GITHUB_PAT")
	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPO")

	if pat == "" || owner == "" || repo == "" {
		log.Fatal("Please set GITHUB_PAT, GITHUB_OWNER, and GITHUB_REPO environment variables")
	}

	// Initialize the issue service
	service := issue.New(issue.Config{
		PATToken: pat,
		Owner:    owner,
		Repo:     repo,
	})

	// Create a new issue
	newIssue := &issue.Issue{
		Title:     "Test Issue from GitHub Issue Tracker",
		Body:      "This is a test issue created using the GitHub Issue Tracker package.",
		Labels:    []string{"test", "documentation"},
		Assignees: []string{owner},
	}

	response, err := service.Create(newIssue)
	if err != nil {
		log.Fatalf("Failed to create issue: %v", err)
	}

	fmt.Printf("Issue created successfully!\n")
	fmt.Printf("Title: %s\n", response.Title)
	fmt.Printf("URL: %s\n", response.HTMLURL)
	fmt.Printf("Number: %d\n", response.Number)

	// Get the created issue
	issue, err := service.Get(response.Number)
	if err != nil {
		log.Fatalf("Failed to get issue: %v", err)
	}

	fmt.Printf("\nRetrieved issue:\n")
	fmt.Printf("Title: %s\n", issue.Title)
	fmt.Printf("State: %s\n", issue.State)

	// Update the issue
	update := &issue.Issue{
		Title: "Updated: Test Issue from GitHub Issue Tracker",
		Body:  "This issue has been updated using the GitHub Issue Tracker package.",
		State: "closed",
	}

	updated, err := service.Update(response.Number, update)
	if err != nil {
		log.Fatalf("Failed to update issue: %v", err)
	}

	fmt.Printf("\nIssue updated successfully!\n")
	fmt.Printf("New title: %s\n", updated.Title)
	fmt.Printf("New state: %s\n", updated.State)
}

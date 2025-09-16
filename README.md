# GitHub Issue Tracker

A Go package for managing GitHub issues securely through your backend services.

## Features

- Create, read, and update GitHub issues
- Secure PAT handling
- Configurable API endpoint
- Custom HTTP client support
- Comprehensive test coverage
- Type-safe API

## Installation

```bash
go get github.com/frans-sjostrom/github-issue-tracker
```

## Quick Start

```go
package main

import (
    "fmt"
    "os"
    "github.com/frans-sjostrom/github-issue-tracker/pkg/issue"
)

func main() {
    // Initialize the service
    service := issue.New(issue.Config{
        PATToken: os.Getenv("GITHUB_PAT"),
        Owner:    "your-username",
        Repo:     "your-repo",
    })

    // Create a new issue
    newIssue := &issue.Issue{
        Title:     "Bug: Application crashes on startup",
        Body:      "Detailed description of the issue...",
        Labels:    []string{"bug", "high-priority"},
        Assignees: []string{"maintainer-username"},
    }

    response, err := service.Create(newIssue)
    if err != nil {
        fmt.Printf("Failed to create issue: %v\n", err)
        return
    }

    fmt.Printf("Issue created: %s\n", response.HTMLURL)
}
```

## API Reference

### Creating an Issue

```go
issue := &issue.Issue{
    Title:     "Feature Request: Dark Mode",
    Body:      "Add dark mode support to the application",
    Labels:    []string{"enhancement"},
    Assignees: []string{"developer1", "developer2"},
    State:     "open",
}

response, err := service.Create(issue)
```

### Getting an Issue

```go
response, err := service.Get(issueNumber)
if err != nil {
    // Handle error
}
fmt.Printf("Issue title: %s\n", response.Title)
```

### Updating an Issue

```go
update := &issue.Issue{
    Title: "Updated: Feature Request: Dark Mode",
    State: "closed",
}

response, err := service.Update(issueNumber, update)
```

## Configuration Options

```go
config := issue.Config{
    PATToken:    "your-github-pat",
    Owner:       "repository-owner",
    Repo:        "repository-name",
    APIEndpoint: "https://api.github.com", // Optional, defaults to GitHub API
    HTTPClient:  customClient,             // Optional, defaults to http.DefaultClient
}
```

## Security Best Practices

1. Never expose GitHub PATs in frontend code
2. Store PATs securely using environment variables or secrets management
3. Implement proper authentication for your API endpoints
4. Validate and sanitize user input before creating issues

## Example Web Service Integration

```go
func CreateIssueHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request
    var req struct {
        Title     string   `json:"title"`
        Body      string   `json:"body"`
        Labels    []string `json:"labels"`
        Assignees []string `json:"assignees"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Create issue using the service
    issueService := issue.New(issue.Config{
        PATToken: os.Getenv("GITHUB_PAT"),
        Owner:    os.Getenv("GITHUB_OWNER"),
        Repo:     os.Getenv("GITHUB_REPO"),
    })

    response, err := issueService.Create(&issue.Issue{
        Title:     req.Title,
        Body:      req.Body,
        Labels:    req.Labels,
        Assignees: req.Assignees,
    })

    if err != nil {
        http.Error(w, "Failed to create issue", http.StatusInternalServerError)
        return
    }

    // Return success response
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Issue created successfully",
        "url":     response.HTMLURL,
    })
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

MIT License - see LICENSE file for details 
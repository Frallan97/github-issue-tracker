package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Service handles GitHub issue operations
type Service struct {
	patToken    string
	owner       string
	repo        string
	apiEndpoint string
	httpClient  *http.Client
}

// Config holds configuration for the issue service
type Config struct {
	PATToken    string
	Owner       string
	Repo        string
	APIEndpoint string
	HTTPClient  *http.Client
}

// Issue represents a GitHub issue
type Issue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Labels    []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
}

// Response represents a GitHub issue response
type Response struct {
	ID      int    `json:"id"`
	Number  int    `json:"number"`
	HTMLURL string `json:"html_url"`
	State   string `json:"state"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	NodeID  string `json:"node_id"`
}

// New creates a new GitHub issue service
func New(config Config) *Service {
	apiEndpoint := config.APIEndpoint
	if apiEndpoint == "" {
		apiEndpoint = "https://api.github.com"
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Service{
		patToken:    config.PATToken,
		owner:       config.Owner,
		repo:        config.Repo,
		apiEndpoint: apiEndpoint,
		httpClient:  httpClient,
	}
}

// Create creates a new GitHub issue
func (s *Service) Create(issue *Issue) (*Response, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues", s.apiEndpoint, s.owner, s.repo)

	payload, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal issue: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.patToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create issue, status: %d", resp.StatusCode)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// Get retrieves a GitHub issue by number
func (s *Service) Get(issueNumber int) (*Response, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d", s.apiEndpoint, s.owner, s.repo, issueNumber)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.patToken))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get issue, status: %d", resp.StatusCode)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// Update updates an existing GitHub issue
func (s *Service) Update(issueNumber int, issue *Issue) (*Response, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d", s.apiEndpoint, s.owner, s.repo, issueNumber)

	payload, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal issue: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.patToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update issue, status: %d", resp.StatusCode)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

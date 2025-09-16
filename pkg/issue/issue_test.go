package issue

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIssue(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/repos/testowner/testrepo/issues", r.URL.Path)
		assert.Equal(t, "Bearer testtoken", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Parse request body
		var issue Issue
		err := json.NewDecoder(r.Body).Decode(&issue)
		assert.NoError(t, err)
		assert.Equal(t, "Test Issue", issue.Title)
		assert.Equal(t, "Test Description", issue.Body)

		// Send response
		response := Response{
			ID:      1,
			Number:  1,
			HTMLURL: "https://github.com/testowner/testrepo/issues/1",
			Title:   issue.Title,
			Body:    issue.Body,
			State:   "open",
			NodeID:  "MDU6SXNzdWUx",
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	// Create service with mock server
	service := New(Config{
		PATToken:    "testtoken",
		Owner:       "testowner",
		Repo:        "testrepo",
		APIEndpoint: server.URL,
	})

	// Create test issue
	issue := &Issue{
		Title: "Test Issue",
		Body:  "Test Description",
	}

	// Test create issue
	response, err := service.Create(issue)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Test Issue", response.Title)
	assert.Equal(t, "https://github.com/testowner/testrepo/issues/1", response.HTMLURL)
}

func TestGetIssue(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/repos/testowner/testrepo/issues/1", r.URL.Path)
		assert.Equal(t, "Bearer testtoken", r.Header.Get("Authorization"))

		// Send response
		response := Response{
			ID:      1,
			Number:  1,
			HTMLURL: "https://github.com/testowner/testrepo/issues/1",
			Title:   "Test Issue",
			Body:    "Test Description",
			State:   "open",
			NodeID:  "MDU6SXNzdWUx",
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create service with mock server
	service := New(Config{
		PATToken:    "testtoken",
		Owner:       "testowner",
		Repo:        "testrepo",
		APIEndpoint: server.URL,
	})

	// Test get issue
	response, err := service.Get(1)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Test Issue", response.Title)
	assert.Equal(t, "https://github.com/testowner/testrepo/issues/1", response.HTMLURL)
}

func TestUpdateIssue(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/repos/testowner/testrepo/issues/1", r.URL.Path)
		assert.Equal(t, "Bearer testtoken", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Parse request body
		var issue Issue
		err := json.NewDecoder(r.Body).Decode(&issue)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Issue", issue.Title)
		assert.Equal(t, "Updated Description", issue.Body)

		// Send response
		response := Response{
			ID:      1,
			Number:  1,
			HTMLURL: "https://github.com/testowner/testrepo/issues/1",
			Title:   issue.Title,
			Body:    issue.Body,
			State:   "open",
			NodeID:  "MDU6SXNzdWUx",
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create service with mock server
	service := New(Config{
		PATToken:    "testtoken",
		Owner:       "testowner",
		Repo:        "testrepo",
		APIEndpoint: server.URL,
	})

	// Create test issue update
	issue := &Issue{
		Title: "Updated Issue",
		Body:  "Updated Description",
	}

	// Test update issue
	response, err := service.Update(1, issue)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Updated Issue", response.Title)
	assert.Equal(t, "https://github.com/testowner/testrepo/issues/1", response.HTMLURL)
}

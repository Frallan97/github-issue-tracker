package issue

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIssue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/repos/testowner/testrepo/issues", r.URL.Path)
		assert.Equal(t, "Bearer testtoken", r.Header.Get("Authorization"))

		var issueReq Issue
		json.NewDecoder(r.Body).Decode(&issueReq)
		assert.Equal(t, "Test Issue", issueReq.Title)

		response := Response{
			ID:      1,
			Number:  1,
			Title:   issueReq.Title,
			HTMLURL: "https://github.com/testowner/testrepo/issues/1",
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	service := New(Config{
		PATToken:    "testtoken",
		Owner:       "testowner",
		Repo:        "testrepo",
		APIEndpoint: server.URL,
	})

	req := &IssueRequest{
		Issue: &Issue{
			Title: "Test Issue",
			Body:  "Test Body",
		},
	}

	response, err := service.Create(req)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Test Issue", response.Title)
	assert.Equal(t, "https://github.com/testowner/testrepo/issues/1", response.HTMLURL)
}

func TestUpdateIssue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/repos/testowner/testrepo/issues/1", r.URL.Path)
		assert.Equal(t, "Bearer testtoken", r.Header.Get("Authorization"))

		var issueReq Issue
		json.NewDecoder(r.Body).Decode(&issueReq)
		assert.Equal(t, "Updated Issue", issueReq.Title)

		response := Response{
			ID:      1,
			Number:  1,
			Title:   issueReq.Title,
			HTMLURL: "https://github.com/testowner/testrepo/issues/1",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	service := New(Config{
		PATToken:    "testtoken",
		Owner:       "testowner",
		Repo:        "testrepo",
		APIEndpoint: server.URL,
	})

	req := &IssueRequest{
		Issue: &Issue{
			Title: "Updated Issue",
			Body:  "Updated Body",
		},
	}

	response, err := service.Update(1, req)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Updated Issue", response.Title)
	assert.Equal(t, "https://github.com/testowner/testrepo/issues/1", response.HTMLURL)
}

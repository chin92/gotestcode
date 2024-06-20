package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Mock the Artifactory response
const mockResponse = `{
	"uri": "http://artifactory.example.com/artifactory/api/storage/my-repo",
	"files": [
		{"uri": "/artifact-1.0.0.jar"},
		{"uri": "/artifact-1.1.0.jar"}
	]
}`

func TestGetLatestArtifactURL(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	expectedURL := "http://artifactory.example.com/artifactory/api/storage/my-repo/artifact-1.1.0.jar"
	actualURL, err := getLatestArtifactURL(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if actualURL != expectedURL {
		t.Fatalf("expected %v, got %v", expectedURL, actualURL)
	}
}

func TestDownloadFile(t *testing.T) {
	// Create a mock server to serve the file
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("dummy content"))
	}))
	defer server.Close()

	filepath := "testfile"
	err := downloadFile(server.URL, filepath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer os.Remove(filepath)

	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		t.Fatalf("expected file to exist, but it does not")
	}
	if info.Size() == 0 {
		t.Fatalf("expected file size to be greater than 0")
	}
}

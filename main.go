package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Artifact represents an artifact in the JFrog Artifactory.
type Artifact struct {
	URI string `json:"uri"`
}

// Response represents the response from Artifactory.
type Response struct {
	URI  string     `json:"uri"`
	Files []Artifact `json:"files"`
}

// getLatestArtifactURL fetches the URL of the latest artifact from the specified Artifactory repository.
func getLatestArtifactURL(repoURL string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(repoURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch artifacts: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if len(response.Files) == 0 {
		return "", fmt.Errorf("no artifacts found")
	}

	latestArtifact := response.Files[len(response.Files)-1]
	return response.URI + latestArtifact.URI, nil
}

// downloadFile downloads a file from the specified URL and saves it to the given path.
func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	repoURL := "http://artifactory.example.com/artifactory/api/storage/my-repo"
	artifactURL, err := getLatestArtifactURL(repoURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filepath := "latest-artifact"
	if err := downloadFile(artifactURL, filepath); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Downloaded latest artifact to", filepath)
	//sed -n 's/.*"uri":"\([^"]*\)".*/\1/p' | sort -V | tail -n 1)
}

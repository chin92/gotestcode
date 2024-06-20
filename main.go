package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// URL of the JSON file in JFrog Artifactory
	url := "https://your-artifactory-instance/artifactory/path/to/your-file.json"
	// Path where the file will be saved locally
	filePath := "local-file.json"

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download file: HTTP Status %s\n", resp.Status)
		return
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error copying the data to file:", err)
		return
	}

	fmt.Println("File downloaded successfully!")
}

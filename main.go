package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	url := "https://example.com/data.json" // Replace with the actual URL
	filepath := "data.json" // Replace with the desired local file path

	err := downloadFile(url, filepath)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
	} else {
		fmt.Println("File downloaded successfully")
	}
}

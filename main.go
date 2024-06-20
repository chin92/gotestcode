package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    // Define the JFrog repository URL and the artifact path
    repoURL := "https://your.jfrog.instance/artifactory/repo-name/path/to/artifact.ext"
    
    // Create a new HTTP request
    req, err := http.NewRequest("GET", repoURL, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
    
    // Set up authentication if needed (Basic Auth example)
    username := "your-username"
    password := "your-password"
    req.SetBasicAuth(username, password)
    
    // Send the HTTP request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()
    
    // Check the response status
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Failed to download file, status code:", resp.StatusCode)
        return
    }
    
    // Create the output file
    outFile, err := os.Create("artifact.ext")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer outFile.Close()
    
    // Copy the response body to the file
    _, err = io.Copy(outFile, resp.Body)
    if err != nil {
        fmt.Println("Error saving file:", err)
        return
    }
    
    fmt.Println("File downloaded successfully")
}

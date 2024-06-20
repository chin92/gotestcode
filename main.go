package main

import (
    "fmt"
    "os"

    "github.com/jfrog/jfrog-client-go/artifactory"
    "github.com/jfrog/jfrog-client-go/artifactory/services"
    "github.com/jfrog/jfrog-client-go/artifactory/services/utils"
    "github.com/jfrog/jfrog-client-go/config"
    "github.com/jfrog/jfrog-client-go/utils/log"
)

func main() {
    // Set up the JFrog configuration
    serverDetails := artifactory.NewConfigBuilder().
        SetUrl("https://your-jfrog-instance.com/artifactory").
        SetUser("your-username").
        SetPassword("your-password").
        Build()

    // Create the Artifactory service manager
    serviceConfig, err := config.NewConfigBuilder().
        SetServiceDetails(serverDetails).
        SetDryRun(false).
        SetThreads(3).
        Build()

    if err != nil {
        log.Error(err)
        return
    }

    serviceManager, err := artifactory.New(&serverDetails, serviceConfig)
    if err != nil {
        log.Error(err)
        return
    }

    // Set up the file download parameters
    params := services.NewDownloadParams()
    params.ArtifactoryCommonParams = &utils.ArtifactoryCommonParams{
        Pattern: "repo/path/to/your/file",
        Target:  "local/path/to/download/file",
    }

    // Download the file
    totalDownloaded, err := serviceManager.DownloadFiles(params)
    if err != nil {
        log.Error(err)
        return
    }

    fmt.Printf("Downloaded %d files\n", totalDownloaded)
}


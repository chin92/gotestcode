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
    // Configure the Artifactory details
    artDetails := auth.NewArtifactoryDetails()
    artDetails.SetUrl("https://your.jfrog.io/artifactory")
    artDetails.SetUser("your-username")
    artDetails.SetPassword("your-password")

    // Create the service manager
    serviceConfig, err := config.NewConfigBuilder().
        SetServiceDetails(artDetails).
        SetDryRun(false).
        Build()
    if err != nil {
        fmt.Println("Failed to create config:", err)
        os.Exit(1)
    }
    sm, err := artifactory.New(serviceConfig)
    if err != nil {
        fmt.Println("Failed to create service manager:", err)
        os.Exit(1)
    }

    // Define the download parameters
    downloadParams := services.DownloadParams{}
    downloadParams.ArtifactoryCommonParams = &utils.ArtifactoryCommonParams{
        Pattern: "my-repo/path/to/file.txt",
        Target:  "./local/path/to/download/",
    }

    // Perform the download
    _, _, failed, err := sm.DownloadFilesWithSummary(downloadParams)
    if err != nil {
        fmt.Println("Failed to download file:", err)
        os.Exit(1)
    }

    if len(failed) > 0 {
        fmt.Println("Download failed for some files:", failed)
    } else {
        fmt.Println("Download successful")
    }
}

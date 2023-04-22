package main

import (
    "fmt"
    "github.com/google/go-github/v35/github"
    "os"
)

func main() {
    // Get the app id, installation id and private key from environment variables.
    appId := os.Getenv("GITHUB_APP_ID")
    installationId := os.Getenv("GITHUB_INSTALLATION_ID")
    privateKey := os.Getenv("GITHUB_PRIVATE_KEY")

    // Create a new GitHub client.
    client := github.NewClient(nil)

    // Create a new app installation token.
    token, err := client.Apps.CreateAppInstallationToken(appId, installationId, &github.CreateAppInstallationTokenOptions{
        PrivateKey: privateKey,
    })
    if err != nil {
        panic(err)
    }

    // Print the token.
    fmt.Println(token.Token)
}

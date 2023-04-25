package main

import (
    "fmt"
    "os"
    "github.com/google/go-github/v35/github"
)

func main() {
    // Get the app ID, installation ID, and private key from environment variables.
    appID := os.Getenv("GITHUB_APP_ID")
    installationID := os.Getenv("GITHUB_INSTALLATION_ID")
    privatePEM := os.Getenv("GITHUB_PRIVATE_KEY")

    // Create a new GitHub client.
    client := github.NewClient(nil)

    // Set the authentication header.
    client.SetAuthorization("Bearer " + privatePEM)

    // Create a new installation access token request.
    body := github.InstallationAccessTokenRequest{
        Name: "My Installation Token",
    }

    // Create a new installation access token.
    accessToken, _, err := client.Apps.CreateInstallationAccessToken(appID, installationID, &body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Print the installation access token.
    fmt.Println(accessToken.Token)
}

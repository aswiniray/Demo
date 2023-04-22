package main

import (
    "context"
    "fmt"
    "github.com/google/go-github/v52/github"
    "os"
)

func main() {
    // Get the app ID, installation ID, and private key from environment variables.
    appID := os.Getenv("GITHUB_APP_ID")
    installationID := os.Getenv("GITHUB_INSTALLATION_ID")
    privatePEMKey := os.Getenv("GITHUB_PRIVATE_KEY_PEM")

    // Create a new GitHub client.
    client := github.NewClient(nil)

    // Create a new JWT transport using the app ID, installation ID, and private key.
    transport, err := github.NewAppsTransportKeyFromFile(client.Transport, appID, installationID, privatePEMKey)
    if err != nil {
        panic(err)
    }

    // Set the transport on the client.
    client.Transport = transport

    // Create a new context.
    ctx := context.Background()

    // Generate a new GitHub access token.
    accessToken, err := client.Apps.CreateAccessToken(ctx, appID, installationID)
    if err != nil {
        panic(err)
    }

    // Print the access token.
    fmt.Println(accessToken.Token)
}

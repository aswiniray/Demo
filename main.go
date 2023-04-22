package main

import (
    "context"
    "fmt"
    "os"
    "github.com/google/go-github/v41/github"
    "github.com/dgrijalva/jwt-go"
)

func main() {
    // Get the app ID, installation ID, and private key from environment variables.
    appID := os.Getenv("GITHUB_APP_ID")
    installationID := os.Getenv("GITHUB_INSTALLATION_ID")
    privateKey := os.Getenv("GITHUB_PRIVATE_KEY")

    // Create a new GitHub client.
    client := github.NewClient(nil)

    // Create a new JWT token with the app ID, installation ID, and private key.
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.Claims{
        "app_id": appID,
        "installation_id": installationID,
    })
    tokenString, err := token.SignedString([]byte(privateKey))
    if err != nil {
        panic(err)
    }

    // Set the authorization header with the JWT token.
    ctx := context.WithValue(context.Background(), github.Authorization, "Bearer "+tokenString)

    // Get the access token from GitHub.
    access, _, err := client.Apps.GetInstallationAccessToken(ctx, appID, installationID)
    if err != nil {
        panic(err)
    }

    // Print the access token.
    fmt.Println(access.Token)
}

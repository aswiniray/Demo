package main

import (
    "fmt"
    "os"
    "github.com/google/go-github/v52/github"
    "github.com/dgrijalva/jwt-go"
)

func main() {
    // Get the app ID, installation ID, and private key from environment variables.
    appID := os.Getenv("GITHUB_APP_ID")
    installationID := os.Getenv("GITHUB_INSTALLATION_ID")
    privateKey := os.Getenv("GITHUB_PRIVATE_KEY")

    // Create a new JWT claims object.
    claims := jwt.NewClaims()
    claims.Set("app_id", appID)
    claims.Set("installation_id", installationID)

    // Create a new JWT signing algorithm.
    signingAlgorithm := jwt.SigningMethodHS256

    // Create a new JWT token.
    token := jwt.NewWithClaims(signingAlgorithm, claims)

    // Sign the JWT token with the private key.
    token.Sign(privateKey)

    // Get the GitHub client.
    client := github.NewClient(nil)

    // Create a new request to create an installation access token.
    request := github.NewCreateInstallationAccessTokenRequest{
        Token: token.Raw,
    }

    // Send the request to create the installation access token.
    response, err := client.Apps.CreateInstallationAccessToken(request)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the installation access token.
    fmt.Println(response.Token)
}


package main

import (
    "context"
    "fmt"
    "github.com/bradleyfalzon/ghinstallation"
    "github.com/google/go-github/v39/github"
    "golang.org/x/oauth2"
)

func main() {
    // Replace with your app ID
    appID := int64(12345)

    // Replace with your installation ID
    installationID := int64(67890)

    // Replace with your private key
    privateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
...
-----END RSA PRIVATE KEY-----`)

    // Create a new transport for authentication
    tr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, appID, installationID, privateKey)
    if err != nil {
        panic(err)
    }

    // Create a new client using the authenticated transport
    client := github.NewClient(&http.Client{Transport: tr})

    // Use the client to create an installation token
    token, _, err := client.Apps.CreateInstallationToken(context.Background(), nil)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Token: %v\n", token.GetToken())
}

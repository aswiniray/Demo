
package main

import (
  "fmt"
  "os"
  "github.com/google/go-github/v32/github"
  "github.com/nabeken/ghinstallation"
)

func main() {
  // Get the app ID, installation ID, and private key from the environment.
  appID := os.Getenv("GITHUB_APP_ID")
  installationID := os.Getenv("GITHUB_INSTALLATION_ID")
  privateKeyPath := os.Getenv("GITHUB_PRIVATE_KEY_PATH")

  // Create a new GitHub client.
  client := github.NewClient(nil)

  // Create a new ghinstallation client.
  ghinstallationClient := ghinstallation.NewClient(client)

  // Create a new installation token request.
  request := ghinstallation.NewCreateInstallationTokenRequest{
    AppID: appID,
     InstallationID: installationID,
  }

  // Create the installation token.
  token, err := ghinstallationClient.CreateInstallationToken(request)
  if err != nil {
    panic(err)
  }

  // Print the installation token.
  fmt.Println(token.Token)
}

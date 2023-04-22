package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

func main() {
	appID := os.Getenv("APP_ID")
	installationID := os.Getenv("INSTALLATION_ID")
	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")

	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)

	ctx := context.Background()
	ts := oauth2.NewClient(ctx, tokenSource).Transport

	client := github.NewClient(&http.Client{Transport: ts})

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat":   jwt.TimeFunc().Unix(),
		"exp":   jwt.TimeFunc().Add(time.Minute * 10).Unix(),
		"iss":   appID,
		"sub":   installationID,
		"nbf":   jwt.TimeFunc().Unix(),
	})

	tokenString, err := jwtToken.SignedString(privateKey)
	if err != nil {
		log.Fatalf("Failed to sign JWT token: %v", err)
	}

	token, _, err := client.Apps.CreateInstallationToken(ctx, installationID, &github.InstallationTokenOptions{
		Permissions: &github.InstallationPermissions{
			Metadata: "read",
			PullRequests: "write",
			Issues: "write",
			Packages: "read",
			Statuses: "read",
			Secrets: "read",
			SingleFile: "write",
			CodeScanningWrite: true,
			CodeScanningRead: true,
			VulnerabilityAlertsRead: true,
			SecurityEventsRead: true,
			MigrationsWrite: true,
			MigrationsRead: true,
			MembersRead: true,
			MembersWrite: true,
			MetadataRead: true,
			PackagesRead: true,
			PagesRead: true,
			PagesWrite: true,
			PullRequestsRead: true,
			IssuesRead: true,
			IssuesWrite: true,
			SecretsWrite: true,
			SecurityEventsWrite: true,
			StatusesWrite: true,
			VulnerabilityAlertsWrite: true,
			GitHubActionsWriteAccess: true,
			GitHubActionsReadAccess: true,
			GitHubPackagesWriteAccess: true,
			GitHubPackagesReadAccess:true
		  },
		  ExpirationTime:"2023-04-23T00:00:00Z",
		  RepositoryIds:nil
	  })
	  if err != nil {
		  log.Fatalf("Failed to create installation token for installation ID %q : %v", installationID, err)
	  }
	  fmt.Printf("Installation token for installation ID %q created successfully.\nToken:\n%s\n", installationID, token.GetToken())
}

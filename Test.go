package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

func main() {
	appID := os.Getenv("APP_ID")
	instID := os.Getenv("INSTALLATION_ID")
	keyPath := os.Getenv("PRIVATE_KEY_PATH")

	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("failed to read private key file: %v", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}

	tr := http.DefaultTransport
	it, err := ghinstallation.NewAppsTransportKeyFromFile(tr, appID, keyPath)
	if err != nil {
		log.Fatalf("failed to create installation transport: %v", err)
	}

	ctx := context.Background()
	ts := oauth2.NewClient(ctx, it).Transport

	client := github.NewClient(&http.Client{Transport: ts})

	token, _, err := client.Apps.CreateInstallationToken(ctx, instID, nil)
	if err != nil {
		log.Fatalf("failed to create installation token: %v", err)
	}

	fmt.Println(token.GetToken())
}

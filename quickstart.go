package main

import (
	"fmt"
	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v2"
	"log"
	"os"
)

// Settings for authorization.
var config = &oauth2.Config{
	ClientID:     "856948250325-v98c9lbbcld6dgnu7ard7fcr41qggvjj.apps.googleusercontent.com",
	ClientSecret: "12LWMqnX-bfvIf8DmeI0Yt0-",
	Scopes:       []string{drive.DriveScope},
	RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
	},
}

// Uploads a file to Google Drive
func main() {

	// Generate a URL to visit for authorization.
	authUrl := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Printf("Go to the following link in your browser: %v\n", authUrl)
	// Read the code, and exchange it for a token.
	log.Printf("Enter verification code: ")
	var code string
	fmt.Scanln(&code)

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("An error occurred exchanging the code: %v\n", err)
	}

	// Create a new authorized Drive client.
	svc, err := drive.New(config.Client(oauth2.NoContext, tok))
	if err != nil {
		log.Fatalf("An error occurred creating Drive client: %v\n", err)
	}

	// Define the metadata for the file we are going to create.
	f := &drive.File{
		Title:       "My Document",
		Description: "My test document",
	}

	// Read the file data that we are going to upload.
	m, err := os.Open("document.txt")
	if err != nil {
		log.Fatalf("An error occurred reading the document: %v\n", err)
	}

	// Make the API request to upload metadata and file data.
	r, err := svc.Files.Insert(f).Media(m).Do()
	if err != nil {
		log.Fatalf("An error occurred uploading the document: %v\n", err)
	}
	log.Printf("Created: ID=%v, Title=%v\n", r.Id, r.Title)
}

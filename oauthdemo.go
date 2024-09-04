// References:
// https://medium.com/@_RR/google-oauth-2-0-and-golang-4cc299f8c1ed
// https://dev.to/siddheshk02/oauth-20-implementation-in-golang-3mj1

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

type User struct {
	Id              string `json:"id"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"verified_email"`
	PictureUrl      string `json:"picture"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occurred loading env variables: %s", err)
	}

	// Initialize the OAuth2 config once during package initialization
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/google-oauth-callback", os.Getenv("BASE_URL")),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func oauthDemo(w http.ResponseWriter, r *http.Request) {
	// TODO - understand how to properly create oauthState to prevent csrf
	// From AuthCodeURL:
	// To protect against CSRF attacks, opts should include a PKCE challenge
	// (S256ChallengeOption). Not all servers support PKCE. An alternative is to
	// generate a random state parameter and verify it after exchange.
	oauthState := ""
	u := googleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use the code to get a token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	// Create a new HTTP request to the userinfo endpoint
	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	// Read the response body
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %s", err.Error())
	}

	// TODO: normally, this is where we would save the user & token
	// saveUser(contents)
	// saveToken(contents, token)

	return contents, nil
}

func getAsHtml(userData []byte) (string, error) {
	var user User
	err := json.Unmarshal(userData, &user)
	if err != nil {
		return "<div>Oauth error!</div>", fmt.Errorf("failed to unmarshal user data: %s", err.Error())
	}

	t, err := template.ParseFiles("./templates/oauth_demo.html")
	if err != nil {
		return "<div>Error executing oauth_demo template!</div>", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, user)
	if err != nil {
		return "<div>Error executing oauth_demo template!</div>", err
	}

	return tpl.String(), nil
}

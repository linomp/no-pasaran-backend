// References:
// https://medium.com/@_RR/google-oauth-2-0-and-golang-4cc299f8c1ed
// https://dev.to/siddheshk02/oauth-20-implementation-in-golang-3mj1

package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	Picture         string `json:"picture"`
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
	defer response.Body.Close()

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
		return "", fmt.Errorf("failed to unmarshal user data: %s", err.Error())
	}

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>My Oauth2 Demo</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 20px;
				padding: 0;
				background-color: #f4f4f4;
			}
			.user-info {
				background-color: #fff;
				border: 1px solid #ddd;
				padding: 20px;
				border-radius: 8px;
				max-width: 400px;
				margin: auto;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			}
			.user-info img {
				max-width: 100px;
				border-radius: 50%%;
				display: block;
				margin-bottom: 20px;
			}
			.user-info .item {
				margin-bottom: 10px;
			}
			.user-info .item label {
				font-weight: bold;
			}
			.user-info .item span {
				margin-left: 10px;
			}
		</style>
	</head>
	<body>

	<div class="user-info">
	    <h1>Oauth2 Demo</h1>	
		<h2>Google sign-in succesful!</h2>
		<div class="item">
			<label>Email:</label>
			<span>%s</span>
		</div>
		<div class="item">
			<label>Google user ID:</label>
			<span>%s</span>
		</div>
		<div class="item">
			<label>Verified Email:</label>
			<span>%t</span>
		</div>
		<div class="item">
			<img src="%s" alt="User Picture">
		</div>
	</div>

	</body>
	</html>
	`, user.Email, user.Id, user.IsEmailVerified, user.Picture)

	return html, nil
}

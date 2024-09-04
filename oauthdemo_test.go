package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGetAsHtml(t *testing.T) {
	user := User{
		Id:              "123456789",
		Email:           "user@example.com",
		IsEmailVerified: true,
		PictureUrl:      "https://example.com/picture.jpg",
	}
	userData, _ := json.Marshal(user)

	html, err := getAsHtml(userData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.Contains(html, user.Id) {
		t.Errorf("HTML output does not contain user ID. Expected %s", user.Id)
	}
	if !strings.Contains(html, user.Email) {
		t.Errorf("HTML output does not contain user Email. Expected %s", user.Email)
	}
	if !strings.Contains(html, user.PictureUrl) {
		t.Errorf("HTML output does not contain user Picture URL. Expected %s", user.PictureUrl)
	}
	if !strings.Contains(html, "true") {
		t.Errorf("HTML output does not contain 'true' for verified email. Expected 'true'")
	}
}

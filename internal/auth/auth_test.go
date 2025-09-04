package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKeyValidHeader(t *testing.T) {
	// arrange
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey key")
	want := "key"
	// act
	got, err := auth.GetAPIKey(headers)
	// assert
	if err != nil {
		t.Fatalf("\nunexpected error:\n%v\n", err)
	}
	if want != got {
		t.Fatalf("\ntest failed.\nwanted: %v\ngot: %v\n", want, got)
	}
}

func TestGetAPIKeyInvalidHeader(t *testing.T) {
	// arrange
	headers := http.Header{}
	headers.Add("hacker", "random")
	// act
	k, err := auth.GetAPIKey(headers)
	// assert
	if err == nil {
		t.Fatalf("\ntest succeeded when should have failed.\nkey:\n%v\n", k)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer abc") // wrong scheme

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}

	headers.Add("Authorization", "ApiKey API_KEY")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal(err)
	}

	if apiKey != "API_KEY" {
		t.Fatal(apiKey)
	}
}

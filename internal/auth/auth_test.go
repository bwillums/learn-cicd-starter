package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var header = http.Header{}
	header.Add("Authorization", "ApiKey Test")
	apiKey, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Error %v", err)
	}

	if apiKey != "Test" {
		t.Errorf("expected: %v, got %v", "Test", apiKey)
	}
}

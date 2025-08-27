package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("should return API key when valid authorization header is provided", func(t *testing.T) {
		// Arrange
		headers := http.Header{}
		expectedAPIKey := "test-api-key-123"
		headers.Set("Authorization", "ApiKey "+expectedAPIKey)

		// Act
		apiKey, err := GetAPIKey(headers)

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if apiKey != expectedAPIKey {
			t.Errorf("expected API key %q, got %q", expectedAPIKey, apiKey)
		}
	})

	t.Run("should return error when no authorization header is provided", func(t *testing.T) {
		// Arrange
		headers := http.Header{}

		// Act
		apiKey, err := GetAPIKey(headers)

		// Assert
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
		if apiKey != "" {
			t.Errorf("expected empty API key, got %q", apiKey)
		}
	})
}

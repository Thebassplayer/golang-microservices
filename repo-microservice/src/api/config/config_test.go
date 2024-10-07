package config

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Load the .env file for testing
	err := godotenv.Load("../../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", secretGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	// Set a test value for the environment variable
	expectedToken := "SECRET_GITHUB_ACCESS_TOKEN"
	os.Setenv(secretGithubAccessToken, expectedToken)

	token := GetGithubAccessToken()
	if token != expectedToken {
		t.Errorf("Expected %s, but got %s", expectedToken, token)
	}

	// Unset the environment variable to test the failure case
	os.Unsetenv(secretGithubAccessToken)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected log.Fatalf to be called, but it was not")
		}
	}()

	GetGithubAccessToken()
}

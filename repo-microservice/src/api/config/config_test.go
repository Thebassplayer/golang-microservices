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
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", secretGithubAccessToken)
}

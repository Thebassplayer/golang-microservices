package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetGithubAccessToken() string {
	// Return the environment variable when needed
	token := os.Getenv(apiGithubAccessToken)
	if token == "" {
		log.Fatalf("Error: GitHub Access Token not set")
	}
	return token
}

package github_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAthorizationHeader(t *testing.T) {
	header := GetAuthorizationHeader("abc123")

	assert.EqualValues(t, "token abc123", header)
}

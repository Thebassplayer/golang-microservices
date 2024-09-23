package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUSerFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "We were expecting 404 when user is not found")
	assert.EqualValues(t, "not_found", err.Code, "We were expecting an error code 'not_found'")
	assert.EqualValues(t, "User 0 not found", err.Message, "We were expecting an error message 'User 0 not found'")

}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "We were not expecting an error when user id is 123")
	assert.NotNil(t, user, "We were expecting a user with id 123")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Raul", user.FirstName)
	assert.EqualValues(t, "Perrito", user.LastName)
	assert.EqualValues(t, "raulperrito@maildeu.com", user.Email)
}

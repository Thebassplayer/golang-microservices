package services

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thebassplayer/golang-microservices/mvc/domain"
	"github.com/thebassplayer/golang-microservices/mvc/utils"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)

}

func TestGetUserNotFoundInDatabase(t *testing.T) {

	// Initialization:
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User 0 not found",
			Code:       "not_found",
		}
	}
	// Execution:
	user, err := UserService.GetUser(0)

	// Validation:
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when user id is 0")
	assert.EqualValues(t, "not_found", err.Code, "We were expecting an error code 'not_found'")
	assert.EqualValues(t, "User 0 not found", err.Message, "We were expecting an error message 'User 0 not found'")
	assert.EqualValues(t, "User 0 not found", err.Message)
}

func TestGetUserNotError(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := UserService.GetUser(123)

	// Validation:
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}

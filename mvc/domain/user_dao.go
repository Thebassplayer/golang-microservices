package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thebassplayer/golang-microservices/mvc/utils"
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type user_dao struct{}

func init() {
	UserDao = &user_dao{}
}

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Raul", LastName: "Perrito", Email: "raulperrito@maildeu.com"},
	}
	UserDao userDaoInterface
)

func (u *user_dao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}

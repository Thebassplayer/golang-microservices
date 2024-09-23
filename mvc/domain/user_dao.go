package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Raul", LastName: "Perrito", Email: "raulperrito@maildeu.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("User %v not found", userId)

}

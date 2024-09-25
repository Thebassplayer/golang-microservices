package services

import (
	"github.com/thebassplayer/golang-microservices/mvc/domain"
	"github.com/thebassplayer/golang-microservices/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}

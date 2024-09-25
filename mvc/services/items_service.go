package services

import (
	"net/http"

	"github.com/thebassplayer/golang-microservices/mvc/domain"
	"github.com/thebassplayer/golang-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItems(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}

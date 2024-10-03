package services

import "github.com/thebassplayer/golang-microservices/src/api/utils/errors"

type repoService struct{}

type repoServiceInterface interface {
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(request interface{}) (interface{}, errors.ApiError) {

}

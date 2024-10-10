package app

import (
	"github.com/thebassplayer/golang-microservices/repo-microservice/src/api/controllers/polo"
	"github.com/thebassplayer/golang-microservices/repo-microservice/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}

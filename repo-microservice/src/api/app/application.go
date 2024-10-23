package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thebassplayer/golang-microservices/repo-microservice/src/api/log"
)

var (
	router *gin.Engine
)

func init() {
	log.Info("Starting mapping URLs", "step:1", "status:pending")
	router = gin.Default()
	log.Info("URLs mapped", "step:2", "status:success")
}

func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}

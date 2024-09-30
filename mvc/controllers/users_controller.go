package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thebassplayer/golang-microservices/mvc/services"
	"github.com/thebassplayer/golang-microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")

	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	log.Printf("About to process user ID: %v", userId)

	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)

}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/thebassplayer/golang-microservices/mvc/services"
	"github.com/thebassplayer/golang-microservices/mvc/utils"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userIdParam := request.URL.Query().Get("user_id")

	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		// Handle the error and return to the client
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(apiErr.StatusCode)
		response.Write([]byte(jsonValue))
		return
	}

	log.Printf("About to process user ID: %v", userId)

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(apiErr.StatusCode)
		response.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonValue)

}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/thebassplayer/golang-microservices/mvc/services"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userIdParam := request.URL.Query().Get("user_id")

	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		// Handle the error and return to the client
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("User Id must be a number"))
		return
	}

	log.Printf("About to process user ID: %v", userId)

	user, err := services.GetUser(userId)

	if err != nil {
		// Handle the error and return to the client
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
		return
	}

	// Return user to the client
	jsonValue, _ := json.Marshal(user)

	response.Write(jsonValue)

}
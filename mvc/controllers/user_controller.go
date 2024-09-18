package controllers

import "net/http"

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello World!"))

}

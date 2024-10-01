package github_provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/thebassplayer/golang-microservices/src/api/clients/restclient"
	"github.com/thebassplayer/golang-microservices/src/api/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func GetAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {

	headers := http.Header{}
	headers.Set(headerAuthorization, GetAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GitHubErrorResponse{
			Status:  strconv.Itoa(http.StatusInternalServerError),
			Message: err.Error(),
		}
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHubErrorResponse{
			Status:  strconv.Itoa(http.StatusInternalServerError),
			Message: "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorResponse{
				Status:  strconv.Itoa(http.StatusInternalServerError),
				Message: "invalid json response body",
			}
		}
		errResponse.Status = strconv.Itoa(response.StatusCode)
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.GitHubErrorResponse{
			Status:  strconv.Itoa(http.StatusInternalServerError),
			Message: "error when trying to unmarshal github create repo response",
		}
	}

	return &result, nil
}

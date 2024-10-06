package services

import (
	"os"
	"testing"

	"github.com/thebassplayer/golang-microservices/repo-microservice/src/api/clients/restclient"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

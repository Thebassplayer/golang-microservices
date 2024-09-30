package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Golang-repo-test",
		Description: "A repo created from a Golang microservice",
		Homepage:    "https://github.com",
		Private:     true,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"Golang-repo-test","description":"A repo created from a Golang microservice","homepage":"https://github.com","private":true}`, string(bytes))
}

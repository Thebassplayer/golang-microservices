package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserJsonSerialization(t *testing.T) {
	user := User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}

	jsonData, err := json.Marshal(user)
	assert.Nil(t, err, "Error marshalling user to JSON")

	expectedJson := `{"id":1,"first_name":"John","last_name":"Doe","email":"john.doe@example.com"}`

	assert.EqualValues(t, expectedJson, string(jsonData), "JSON marshalling is not as expected")
}

func TestUserJsonDeserialization(t *testing.T) {
	jsonData := `{"id":1,"first_name":"John","last_name":"Doe","email":"john.doe@example.com"}`
	var user User

	err := json.Unmarshal([]byte(jsonData), &user)
	assert.Nil(t, err, "Error unmarshalling JSON to user")

	expectedUser := User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}

	assert.EqualValues(t, expectedUser, user, "User unmarshalling is not as expected")
}

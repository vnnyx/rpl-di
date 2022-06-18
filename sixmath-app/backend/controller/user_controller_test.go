package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"rpl-sixmath/model"
	"testing"
)

func TestUserController_CreateStudent(t *testing.T) {
	user, _ := userRepository.FindAllTeacher()
	var userId int
	for _, users := range user {
		userId = users.UserId
	}
	userRepository.DeleteUser(userId)
	createStudentRequest := model.StudentCreateRequest{
		Email:     "email@email.com",
		Username:  "user1",
		Handphone: "081234567890",
		Password:  "password",
	}
	requestBody, _ := json.Marshal(createStudentRequest)

	request := httptest.NewRequest("POST", "/api/user/student", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createStudentResponse := model.StudentCreateResponse{}
	json.Unmarshal(jsonData, &createStudentResponse)
	assert.NotNil(t, createStudentResponse.UserId)
	assert.Equal(t, createStudentRequest.Username, createStudentResponse.Username)
	assert.Equal(t, createStudentRequest.Email, createStudentResponse.Email)
}

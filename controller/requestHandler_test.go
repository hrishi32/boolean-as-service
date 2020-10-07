package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/hrishi32/boolean-as-service/models"

	"github.com/google/uuid"

	"github.com/hrishi32/boolean-as-service/mocks"

	"github.com/golang/mock/gomock"
)

// Sort out all errors properly, tests will depend on this only

// Get
func TestGetSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	expectedBoolean := models.Boolean{
		ID:    demoUUID,
		Value: true,
		Key:   "somekey",
	}

	mockRepo.EXPECT().Get(demoUUID).Return(expectedBoolean, nil)

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.GET("/:id", GetHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodGet, "/"+demoUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, http.StatusOK, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}
	responseBoolean := models.Boolean{}
	err = json.Unmarshal(responseBody, &responseBoolean)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedBoolean, responseBoolean)

}
func TestGet400(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	badUUID := "A_Bad_UUID"

	// no need to mock Get function as it will not be called in this case.

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.GET("/:id", GetHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodGet, "/"+badUUID, nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, http.StatusBadRequest, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	a := []byte{}
	assert.Equal(t, a, responseBody)
}
func TestGet404(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	expectedBoolean := models.Boolean{}

	mockRepo.EXPECT().Get(demoUUID).Return(expectedBoolean, errors.New("Record not found"))

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.GET("/:id", GetHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodGet, "/"+demoUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, 404, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	a := []byte{}
	assert.Equal(t, a, responseBody)
}
func TestGet500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	expectedBoolean := models.Boolean{}

	mockRepo.EXPECT().Get(demoUUID).Return(expectedBoolean, errors.New("Some new error"))

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.GET("/:id", GetHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodGet, "/"+demoUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	a := []byte{}
	assert.Equal(t, a, responseBody)
}

/* **************************************** */

// Post Tests

func TestPostSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	demoBoolean := models.Boolean{
		Value: true,
		Key:   "demo key",
	}

	expectedBoolean := models.Boolean{
		ID:    demoUUID,
		Value: demoBoolean.Value,
		Key:   demoBoolean.Key,
	}
	mockRepo.EXPECT().Create(demoBoolean).Return(demoUUID, nil)

	// Preservice
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.POST("/", PostHandler)
	server.Run(":8000")

	// Make request

	requestBody := strings.NewReader(`{
		"key": "demo key",
		"value": true
	  }`)
	request, err := http.NewRequest(http.MethodPost, "/", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response
	// Response Code verification
	assert.Equal(t, http.StatusOK, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)
	responseBoolean := models.Boolean{}

	err = json.Unmarshal(responseBody, &responseBoolean)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedBoolean, responseBoolean)

}
func TestPost400(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	// badUUID := "A_Bad_UUID"
	badRequestBody := strings.NewReader(`{
		"key": false,
		"value": 22
	  }`)

	// no need to mock Get function as it will not be called in this case.

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.POST("/", PostHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodPost, "/", badRequestBody)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, http.StatusBadRequest, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	expectedResponse := []byte{}
	assert.Equal(t, expectedResponse, responseBody)
}

// POST 404 does not exist
func TestPost404(t *testing.T) {}
func TestPost500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.UUID{}
	demoBoolean := models.Boolean{
		Value: true,
		Key:   "demo key",
	}

	// expectedBoolean := models.Boolean{}
	mockRepo.EXPECT().Create(demoBoolean).Return(demoUUID, errors.New("Some new error"))

	// Preservice
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.POST("/", PostHandler)
	server.Run(":8000")

	// Make request

	requestBody := strings.NewReader(`{
		"key": "demo key",
		"value": true
	  }`)
	request, err := http.NewRequest(http.MethodPost, "/", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response
	// Response Code verification
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)
	// responseBoolean := models.Boolean{}

	// err = json.Unmarshal(responseBody, &responseBoolean)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	expectedResponse := []byte{}
	assert.Equal(t, expectedResponse, responseBody)
}

// PATCH Tests
func TestPatchSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	// We update a boolean to following key and value, we don't care about previous values
	demoBoolean := models.Boolean{
		Value: true,
		Key:   "demo key",
	}

	expectedBoolean := models.Boolean{
		ID:    demoUUID,
		Value: demoBoolean.Value,
		Key:   demoBoolean.Key,
	}
	mockRepo.EXPECT().Update(demoUUID, demoBoolean).Return(nil)

	// Preservice
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.PATCH("/:id", PatchHandler)
	server.Run(":8000")

	// Make request

	requestBody := strings.NewReader(`{
		"key": "demo key",
		"value": true
	  }`)
	request, err := http.NewRequest(http.MethodPatch, "/"+demoUUID.String(), requestBody)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response
	// Response Code verification
	assert.Equal(t, http.StatusOK, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)
	responseBoolean := models.Boolean{}

	err = json.Unmarshal(responseBody, &responseBoolean)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedBoolean, responseBoolean)
}
func TestPatch400(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	badUUID := "A_Bad_UUID"
	demoUUID := uuid.New()

	badRequestBody := strings.NewReader(`{
		"key": false,
		"value": 22
	  }`)

	requestBody := strings.NewReader(`{
		"key": "demo key",
		"value": true
	  }`)
	// no need to mock Get function as it will not be called in this case.

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.PATCH("/:id", PatchHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request1, err1 := http.NewRequest(http.MethodPatch, "/"+badUUID, requestBody)
	if err1 != nil {
		t.Fatal(err1)
	}

	response1 := httptest.NewRecorder()
	server.ServeHTTP(response1, request1)

	// Check response from server

	assert.Equal(t, http.StatusBadRequest, response1.Code)

	responseBody1, err1 := ioutil.ReadAll(response1.Body)

	if err1 != nil {
		t.Fatal(err1)
	}

	// Checking whether we have empty response
	expectedResponse1 := []byte{}
	assert.Equal(t, expectedResponse1, responseBody1)

	// Another request

	request2, err2 := http.NewRequest(http.MethodPatch, "/"+demoUUID.String(), badRequestBody)
	if err2 != nil {
		t.Fatal(err2)
	}

	response2 := httptest.NewRecorder()
	server.ServeHTTP(response2, request2)

	// Check response from server

	assert.Equal(t, http.StatusBadRequest, response2.Code)

	responseBody2, err2 := ioutil.ReadAll(response2.Body)

	if err2 != nil {
		t.Fatal(err2)
	}

	// Checking whether we have empty response
	expectedResponse2 := []byte{}
	assert.Equal(t, expectedResponse2, responseBody2)

}
func TestPatch404(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()

	requestBody := strings.NewReader(`{
		"key": "demo key",
		"value": true
	  }`)
	// expectedBoolean := models.Boolean{}

	mockRepo.EXPECT().Update(demoUUID, gomock.Any()).Return(errors.New("Record not found"))

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.PATCH("/:id", PatchHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodPatch, "/"+demoUUID.String(), requestBody)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, 404, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	a := []byte{}
	assert.Equal(t, a, responseBody)
}
func TestPatch500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	demoBoolean := models.Boolean{
		Value: true,
		Key:   "demo key",
	}

	// expectedBoolean := models.Boolean{}
	mockRepo.EXPECT().Update(demoUUID, demoBoolean).Return(errors.New("Some new error"))

	// Preservice
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.PATCH("/:id", PatchHandler)
	server.Run(":8000")

	// Make request

	requestBody := strings.NewReader(`{
		"key": "demo key",
		"value": true
	  }`)
	request, err := http.NewRequest(http.MethodPatch, "/"+demoUUID.String(), requestBody)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response
	// Response Code verification
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)
	// responseBoolean := models.Boolean{}

	// err = json.Unmarshal(responseBody, &responseBoolean)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	expectedResponse := []byte{}
	assert.Equal(t, expectedResponse, responseBody)
}

func TestDeleteSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	// expectedBoolean := models.Boolean{
	// 	ID:    demoUUID,
	// 	Value: true,
	// 	Key:   "somekey",
	// }

	mockRepo.EXPECT().Delete(demoUUID).Return(nil)

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.DELETE("/:id", DeleteHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodDelete, "/"+demoUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, http.StatusNoContent, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}
	// responseBoolean := models.Boolean{}
	// err = json.Unmarshal(responseBody, &responseBoolean)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	expectedResponse := []byte{}
	assert.Equal(t, expectedResponse, responseBody)

}
func TestDelete400(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	badUUID := "A_Bad_UUID"

	// no need to mock Get function as it will not be called in this case.

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.DELETE("/:id", DeleteHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodDelete, "/"+badUUID, nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, http.StatusBadRequest, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	a := []byte{}
	assert.Equal(t, a, responseBody)
}
func TestDelete404(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	// expectedBoolean := models.Boolean{}

	mockRepo.EXPECT().Delete(demoUUID).Return(errors.New("Record not found"))

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.DELETE("/:id", DeleteHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodDelete, "/"+demoUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, 404, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	a := []byte{}
	assert.Equal(t, a, responseBody)

}
func TestDelete500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepo(ctrl)

	demoUUID := uuid.New()
	// expectedBoolean := models.Boolean{}

	mockRepo.EXPECT().Delete(demoUUID).Return(errors.New("Some new error"))

	// Preservice setup
	models.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	server.DELETE("/:id", DeleteHandler)
	server.Run(":8000")

	// Make request to abobe created server
	request, err := http.NewRequest(http.MethodDelete, "/"+demoUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	// Check response from server

	// Response Code verification
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	// Response Body verification
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	// Checking whether we have empty response
	a := []byte{}
	assert.Equal(t, a, responseBody)
}

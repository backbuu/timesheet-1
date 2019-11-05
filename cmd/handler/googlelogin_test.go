package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"timesheet/cmd/mockapi"

	. "timesheet/cmd/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_OauthGoogleLogout_Input_AccessToken_Should_Be_304(t *testing.T) {
	request := httptest.NewRequest("POST", "/logout", nil)
	writer := httptest.NewRecorder()
	request.Header.Add("Authorization", "Bearer ya29.Il-vB2mB0hkAEN8KdupS3ZEaXBOHk6qhVntGSkeyAMz6KEoJOpwhfHHQF2KT9W2oiwE1op4pZiUuebKcQ1SBRgRlxMRJxB6Qjf0tl86C5Jdsf51thN-yqvZDBUmUx3hnqw")
	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("DeleteAuthentication", mock.Anything).Return(nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}
	testRoute := gin.Default()
	testRoute.POST("/logout", api.OauthGoogleLogout)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()

	assert.Equal(t, http.StatusTemporaryRedirect, response.StatusCode)
}

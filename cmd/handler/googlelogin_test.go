package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "timesheet/cmd/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_OauthGoogleLogout_Input_AccessToken_Should_Be_200(t *testing.T) {
	request := httptest.NewRequest("POST", "/logout", nil)
	writer := httptest.NewRecorder()
	request.Header.Add("Authorization", "eyJhbGciOiJSUzI1NiIsImtpZCI6ImEwNjgyNGI3OWUzOTgyMzk0ZDVjZTdhYzc1YmY5MmNiYTMwYTJlMjUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI2OTI1NzU4OTgzOTctZG50OXNxaTJqc3RkZGZlcHNuZzA0cDlhYzRvajdwNG4uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI2OTI1NzU4OTgzOTctZG50OXNxaTJqc3RkZGZlcHNuZzA0cDlhYzRvajdwNG4uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTAzMDYxODkyODYyMDM5OTgxMzIiLCJlbWFpbCI6ImxvZ2ludGVzdDUzNUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6InRlWmZfdnZoVTQxQXBqTWdxbGFvX1EiLCJpYXQiOjE1NzM0NjIxNzQsImV4cCI6MTU3MzQ2NTc3NH0.FieIq3nqnEk4sKgNN3gOAHRat-Gj7ewvLV6ri9P4k1_PsoBOSL2brb02HAYrYFYl1NPFwymcp96j_5ZbZnV2k2JbhXvaocPc75pUO8pfzNzVzSp8JiU-OpqUb5CSoguJ6ejLTTGLzFkZ2Uu51GY0Kb_SNkSMGXHwIOlIdSx2UzqrfAqZAliSp_5D1Cp7Ot1I95uv0C79h3TB0ODY9zESsP4lF542ic9sseCt7KCfmoh9hq24OBW9nRLOPqXhOgInvvtqghQd2p7nv88GUdMuCOAFJZgg3_5zoLPkGBiAJcdwwcCoU-kd6r6mcxjKN2xbwFa4G5NskLzNRpUlJQpSRA")

	api := TimesheetAPI{}
	testRoute := gin.Default()
	testRoute.POST("/logout", api.OauthGoogleLogout)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

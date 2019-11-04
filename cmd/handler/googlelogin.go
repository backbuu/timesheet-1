package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	HD            string `json:"hd"`
}

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func OauthGoogleLogin(c *gin.Context) {
	oauthState := generateStateOauthCookie(c.Writer)
	url := googleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func OauthGoogleCallback(context *gin.Context) {
	oauthState, _ := context.Request.Cookie("oauthstate")
	if context.Request.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		context.Redirect(http.StatusTemporaryRedirect, "/home")
		return
	}
	token, _ := getToken(context.Request.FormValue("code"))
	bearer := "Bearer " + token.AccessToken
	context.Writer.Header().Set("Authorization", bearer)

	// data, err := getUserDataFromGoogle(context.Request.FormValue("code"))
	// if err != nil {
	// 	log.Println(err.Error())
	// 	context.Redirect(http.StatusTemporaryRedirect, "/home")
	// 	return
	// }
	// var userInfo UserInfo
	// json.Unmarshal(data, &userInfo)
	context.Redirect(http.StatusTemporaryRedirect, "/home")
	// context.JSON(http.StatusOK, userInfo)
}

func SendAccassToken(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "/home")
}

func generateStateOauthCookie(writer http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	bytes := make([]byte, 16)
	rand.Read(bytes)
	state := base64.URLEncoding.EncodeToString(bytes)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(writer, &cookie)
	return state
}

func getToken(code string) (*oauth2.Token, error) {
	return googleOauthConfig.Exchange(context.Background(), code)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

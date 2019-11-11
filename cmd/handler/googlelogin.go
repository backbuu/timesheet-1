package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"time"
	"timesheet/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

type RefreshTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

func OauthGoogleLogin(context *gin.Context) {
	oauthState := generateStateOauthCookie(context.Writer)
	url := googleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	context.Redirect(http.StatusTemporaryRedirect, url)
}

func (api TimesheetAPI) OauthGoogleLogout(context *gin.Context) {
	requestToken := context.GetHeader("Authorization")
	splitToken := strings.Split(requestToken, "Bearer ")
	requestToken = splitToken[1]
	err := api.TimesheetRepository.DeleteAuthentication(requestToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cookie := &http.Cookie{Name: "access_token", Value: "", Path: "/", Expires: time.Unix(0, 0), HttpOnly: true}
	http.SetCookie(context.Writer, cookie)
	context.Status(http.StatusOK)
}

func DeleteOauthStateCookie(context *gin.Context) {
	cookie := &http.Cookie{Name: "oauthstate", Value: "", Path: "/", Expires: time.Unix(0, 0)}
	http.SetCookie(context.Writer, cookie)
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) OauthGoogleCallback(context *gin.Context) {
	oauthState, _ := context.Request.Cookie("oauthstate")
	if context.Request.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		context.Redirect(http.StatusTemporaryRedirect, "/home")
		return
	}
	token, err := getToken(context.Request.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		context.Redirect(http.StatusTemporaryRedirect, "/home")
		return
	}
	log.Panicln(token)
	userInfo, err := getUserDataFromGoogle(token.AccessToken)
	if err != nil {
		log.Println(err.Error())
		context.Redirect(http.StatusTemporaryRedirect, "/home")
		return
	}
	err = api.TimesheetRepository.CreateAuthentication(userInfo, model.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cookie := http.Cookie{Name: "access_token", Value: token.AccessToken, Expires: time.Now().Add(365 * 24 * time.Hour)}
	http.SetCookie(context.Writer, &cookie)
	context.Redirect(http.StatusTemporaryRedirect, "/home")
}

func generateStateOauthCookie(writer http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err.Error())
	}
	state := base64.URLEncoding.EncodeToString(bytes)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(writer, &cookie)
	return state
}

func getToken(code string) (*oauth2.Token, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	return token, nil
}

func getUserDataFromGoogle(accessToken string) (model.UserInfo, error) {
	var userInfo model.UserInfo
	response, err := http.Get(oauthGoogleUrlAPI + accessToken)
	if err != nil {
		return userInfo, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return userInfo, fmt.Errorf("failed read response: %s", err.Error())
	}
	err = json.Unmarshal(contents, &userInfo)
	if err != nil {
		return userInfo, fmt.Errorf("failed unmashal response: %s", err.Error())
	}
	return userInfo, nil
}

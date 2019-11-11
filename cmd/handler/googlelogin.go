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

	"time"
	"timesheet/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
const daysInYear = 365
const hoursInDay = 24

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func OauthGoogleLogin(context *gin.Context) {
	oauthState := generateStateOauthCookie(context.Writer)
	url := googleOauthConfig.AuthCodeURL(oauthState)
	context.Redirect(http.StatusTemporaryRedirect, url)
}

func (api TimesheetAPI) OauthGoogleLogout(context *gin.Context) {
	cookie := &http.Cookie{Name: "id_token", Value: "", Path: "/", Expires: time.Unix(0, 0), HttpOnly: true}
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
		context.Redirect(http.StatusInternalServerError, "/home?error=invalid_oauth_google_state")
		return
	}
	token, err := getToken(context.Request.FormValue("code"))
	if err != nil {
		context.Redirect(http.StatusInternalServerError, "/home?error=code_exchange_wrong"+err.Error())
		return
	}
	userInfo, err := getUserDataFromGoogle(token.AccessToken)
	if err != nil {
		context.Redirect(http.StatusInternalServerError, "/home?error="+err.Error())
		return
	}
	err = api.Repository.UpdatePictureToMembers(userInfo.Picture, userInfo.Email)
	if err != nil {
		context.Redirect(http.StatusInternalServerError, "/home?error="+err.Error())
		return
	}
	var expiration = now().Add(daysInYear * hoursInDay * time.Hour)
	cookie := http.Cookie{Name: "id_token", Value: token.Extra("id_token").(string), Expires: expiration}
	http.SetCookie(context.Writer, &cookie)
	context.Redirect(http.StatusTemporaryRedirect, "/home")
}

func generateStateOauthCookie(writer http.ResponseWriter) string {
	var expiration = now().Add(daysInYear * hoursInDay * time.Hour)
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
	return googleOauthConfig.Exchange(context.Background(), code)
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

func now() time.Time {
	if os.Getenv("FIX_TIME") != "" {
		fixedTime, _ := time.Parse("20060102150405", os.Getenv("FIX_TIME"))
		return fixedTime
	}
	return time.Now()
}

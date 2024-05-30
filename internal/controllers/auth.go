package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"sonalsguild/helpers"
	"strings"
)

// LOGIN/AUTHENTICATE USER
func GetAutheniticateSpotify(w http.ResponseWriter, req *http.Request) {
	randomString := helpers.GenerateRandomString(16)
	scope := os.Getenv("SPOTIFY_SCOPE")
	redirect_uri := os.Getenv("SPOTIFY_REDIRECT")
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")

	urls := 
	"https://accounts.spotify.com/authorize?client_id=" + client_id + 
	"&redirect_uri=" + redirect_uri +
	"&response_type=code" +
	"&state=" + randomString +
	"&scope=" + strings.ReplaceAll(scope, " ", "+")


	helpers.WriteJson(w, http.StatusOK, urls)
}

// Callback
func GetAuthCallbackSpotify(w http.ResponseWriter, r *http.Request) {
	url := "https://accounts.spotify.com/api/token"
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	spotify_secret := os.Getenv("SPOTIFY_SECRET")
	redirect_uri := os.Getenv("SPOTIFY_REDIRECT")


	if state == "" {
		helpers.WriteJson(w, http.StatusBadRequest, "Invalid query")
	}

	// Prepare the authorization header
	authHeader := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + spotify_secret))
	reqBody := fmt.Sprintf("code=%s&redirect_uri=%s&grant_type=authorization_code", code, redirect_uri)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic " + authHeader)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err!= nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	helpers.WriteJson(w, http.StatusOK, body)

	fmt.Println(string(body))
	// http.Redirect(w,r,"http://localhost:3000/dashboard",http.StatusFound)
}


func GetAuthLogoutSpotify(w http.ResponseWriter, r *http.Request){
	fmt.Println(":sdsd")
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}


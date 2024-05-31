package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sonalsguild/helpers"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func generateJWT(spotify_token AuthResponse ) (string, error) {
	// Define the claims for the JWT
	claims := jwt.MapClaims{
		"access_token":     spotify_token.Access_token, // Replace with your actual access token
		"refresh_token":    spotify_token.Refresh_token, // Replace with your actual refresh token
		"spotify_expires_in":spotify_token.Expires_in,                 // Expires in 1 hour (3600 seconds)
		"issued_at":        time.Now().Unix(),
		"jwt_expires_in":  time.Now().Add(time.Hour * 1).Unix(), // 1 hour from now
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	
	if err!= nil {
		return "", err
	}

	return tokenString, nil
}


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

	// Read the response body and ensure it does not have any errors
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode JSON response
	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)

	if err!= nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	token, err := generateJWT(authResp)

	if err!= nil {
		fmt.Println("Error generating Token:", err)
		return
	}

	fmt.Println("token",token)

	helpers.WriteJson(w, http.StatusOK, token)
}


func GetAuthLogoutSpotify(w http.ResponseWriter, r *http.Request){
	fmt.Println(":sdsd")
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}


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

type AuthResponse struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in  int    `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func generateJWT(access_token string, refresh_token string, expires_in int, ) (string, error) {
	// Define the claims for the JWT
	claims := jwt.MapClaims{
		"access_token":     "your-access-token", // Replace with your actual access token
		"refresh_token":    "your-refresh-token", // Replace with your actual refresh token
		"expires_in":       3600 * 2,                 // Expires in 1 hour (3600 seconds)
		"issued_at":        time.Now().Unix(),
		"expiration_time":  time.Now().Add(time.Hour * 1).Unix(), // 1 hour from now
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err!= nil {
		return "", err
	}

	return tokenString, nil
}


// func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
//  	return 	
// }

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
	body, _ := io.ReadAll(resp.Body)
	if err!= nil {
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

	fmt.Println("Access Token:", authResp.Access_token)
	fmt.Println("Token Type:", authResp.Token_type)
	fmt.Println("Expires In:", authResp.Expires_in)
	fmt.Println("Refresh Token:", authResp.Refresh_token)
	fmt.Println("Scope:", authResp.Scope)

	helpers.WriteJson(w, http.StatusOK, body)

	// fmt.Println(string(body))
	// http.Redirect(w,r,"http://localhost:3000/dashboard",http.StatusFound)
}


func GetAuthLogoutSpotify(w http.ResponseWriter, r *http.Request){
	fmt.Println(":sdsd")
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}


package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sonalsguild/helpers"
	"sonalsguild/internal/services"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var authService services.Authorization

func generateJWT(spotify_token AuthResponse, email string, display_name string , spotify_id string) (string, error) {
	// Define the claims for the JWT
	claims := jwt.MapClaims{
		"access_token":     spotify_token.Access_token, // Replace with your actual access token
		"refresh_token":    spotify_token.Refresh_token, // Replace with your actual refresh token
		"spotify_expires_in":spotify_token.Expires_in,                 // Expires in 1 hour (3600 seconds)
		"issued_at":        time.Now().Unix(),
		"jwt_expires_in":  time.Now().Add(time.Hour * 1).Unix(), // 1 hour from now
		"email":	email,
		"display_name": display_name,
		"spotify_id" : spotify_id,
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
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	authResp, err := RetrieveSpotifyAuthorizationToken(state, code)
	
	if err != nil{
		fmt.Println("Error retrieving token from spotify:", err)
		return
	}

	userResp, err := GetSpotifyUser(*authResp)

	if err != nil{
		fmt.Println("Error retrieving user from spotify:", err)
		return
	}

	exists := authService.CheckUserExists(userResp.Email);
	// Register 
	if(!exists) {

		_,err := authService.RegisterUser(userResp.ID,userResp.DisplayName,userResp.Email)
		if err != nil {
			fmt.Println("Error registering user",err)
			helpers.ErrorJson(w,err)
			return 
		}
	}


	token, err := generateJWT(*authResp,userResp.Email,userResp.DisplayName, userResp.ID)

	if err != nil {
		fmt.Println("Error generating Token:", err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, helpers.Envelop{"authToken":token,"spotify":&authResp})
}

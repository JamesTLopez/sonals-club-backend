package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)



func GetSpotifyUser(auth AuthResponse) (*GetUserSpotify, error) {
	// create get user request
	req, err := http.NewRequest("GET","https://api.spotify.com/v1/me",nil)

	if err != nil {
		fmt.Println("Error decoding generating new request", err)
		return nil, err
	}
	
	// Add headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + auth.Access_token)

	// create http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err!= nil {
		fmt.Println("Error retrieving user information token:", err)
		return nil, err
	}
	
	defer resp.Body.Close()
	getBody, err := io.ReadAll(resp.Body)

	if err!= nil {
		fmt.Println("Error retrieving user information token:", err)
		log.Fatal(err)
		return nil, err
	}
	// Decode JSON response
	var userRepsonse GetUserSpotify
	err = json.Unmarshal(getBody, &userRepsonse)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}


	return &userRepsonse, nil
}

func RetrieveSpotifyAuthorizationToken(state string, code string) (*AuthResponse, error) {
	url := "https://accounts.spotify.com/api/token"
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	spotify_secret := os.Getenv("SPOTIFY_SECRET")
	redirect_uri := os.Getenv("SPOTIFY_REDIRECT")



	// Prepare the authorization header
	authHeader := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + spotify_secret))
	reqBody := fmt.Sprintf("code=%s&redirect_uri=%s&grant_type=authorization_code", code, redirect_uri)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(reqBody))
	if err != nil {
		fmt.Println("Error generating new request", err)
		return nil,err
	}
	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic " + authHeader)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err!= nil {
		fmt.Println("Error retrieving user information token:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body and ensure it does not have any errors
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error retrieving user information token:", err)
		return nil, err
	}

	// Decode JSON response
	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &authResp,nil
}

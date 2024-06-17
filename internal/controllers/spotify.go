package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sonalsguild/helpers"

	"github.com/golang-jwt/jwt/v5"
)


func GetNewReleases(w http.ResponseWriter, r *http.Request) {
	userValue := r.Context().Value("user").(jwt.MapClaims)
	access_token, ok := userValue["access_token"].(string)

	if !ok {
		helpers.MessageLogs.ErrorLog.Println("Something went wrong when grabbing the id")

		return
	}

	response ,err := helpers.GenerateSpotifyGetRequest(w,r,access_token,"browse/new-releases?limit=2")
	
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Something went wrong making the request to spotify")
		return 
	}


	var newReleasesResponse AlbumsResponse
	err = json.Unmarshal(*response, &newReleasesResponse)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Something went wrong marshalling content")
		return 
	}

	
	helpers.WriteJson(w, http.StatusOK, &newReleasesResponse)
}



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


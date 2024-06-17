package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GenerateSpotifyRequest(w http.ResponseWriter, r *http.Request,authToken string, url string) (*[]byte,error) {

	req, err := http.NewRequest("GET","https://api.spotify.com/v1/" + "url",nil)

	if err != nil {
		fmt.Println("Error decoding generating new request", err)
		return nil, err
	}
	
	// Add headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + authToken)

	// create http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err!= nil {
		fmt.Println("Error retrieving user information token:", err)
		return  nil,err
	}
	
	defer resp.Body.Close()
	getBody, err := io.ReadAll(resp.Body)

	if err!= nil {
		fmt.Println("Error retrieving user information token:", err)
		log.Fatal(err)
		return nil, err
	}

	return &getBody,nil
}

package controllers

type AuthResponse struct {
	Access_token 	string `json:"access_token"`
	Token_type   	string `json:"token_type"`
	Expires_in  	int    `json:"expires_in"`
	Refresh_token 	string `json:"refresh_token"`
	Scope        	string `json:"scope"`
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ID              string `json:"id"`
}

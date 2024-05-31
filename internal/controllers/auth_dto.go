package controllers

type AuthResponse struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in  int    `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

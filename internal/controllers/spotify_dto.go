package controllers

// Root represents the top-level structure of the JSON.
type GetUserSpotify struct {
	Country          string                 `json:"country"`
	DisplayName      string                 `json:"display_name"`
	Email            string                 `json:"email"`
	ExplicitContent  ExplicitContent       `json:"explicit_content"`
	ExternalUrls     ExternalUrls          `json:"external_urls"`
	Followers        Followers             `json:"followers"`
	Href             string                 `json:"href"`
	ID               string                 `json:"id"`
	Images           []Image               `json:"images"`
	Product          string                 `json:"product"`
	Type             string                 `json:"type"`
	URI              string                 `json:"uri"`
}

// ExplicitContent represents the explicit content part of the JSON.
type ExplicitContent struct {
	FilterEnabled bool   `json:"filter_enabled"`
	FilterLocked  bool   `json:"filter_locked"`
}

// ExternalUrls represents the external URLs part of the JSON.
type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

// Followers represents the followers part of the JSON.
type Followers struct {
	Href string `json:"href"`
	Total int    `json:"total"`
}

// Image represents individual images within the images array.
type Image struct {
	URL      string `json:"url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}
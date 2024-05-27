package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

// GET Spotify Login callback
func GetAuthCallbackSpotify(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r ,"provider")
	r = r.WithContext(context.WithValue(context.Background(),"provider",provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user)
	http.Redirect(w,r,"http://localhost:3000",http.StatusFound)
}


func GetAuthLogoutSpotify(w http.ResponseWriter, r *http.Request){
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func GetReAutheniticateSpotify(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(res, req); err == nil {

	} else {
		gothic.BeginAuthHandler(res, req)
	}
}
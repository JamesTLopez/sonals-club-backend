package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func GetAuthCallbackSpotify(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r ,"provider")
	r = r.WithContext(context.WithValue(context.Background(),"provider",provider))
	fmt.Println("testinggg")

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println("testingggsss",err)

		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user)
	http.Redirect(w,r,"http://localhost:3000/dashboard",http.StatusFound)
}


func GetAuthLogoutSpotify(w http.ResponseWriter, r *http.Request){
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func GetReAutheniticateSpotify(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r ,"provider")
	r = r.WithContext(context.WithValue(context.Background(),"provider",provider))
	 _, err := gothic.CompleteUserAuth(w, r);
	
	if err == nil {
		return
	}
	gothic.BeginAuthHandler(w, r)

}
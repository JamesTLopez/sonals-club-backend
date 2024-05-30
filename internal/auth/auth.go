package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/spotify"
)



const (
	key = "sdjasdmasd"
	maxAge = 600000
	isProd = false
)

func NewAuth () {
	enverr := godotenv.Load(".env")
	if enverr != nil{
  		log.Fatalf("Error loading .env file in Authfile: %s", enverr)
 	}
	
	 store := sessions.NewCookieStore([]byte(key))
	 store.MaxAge(maxAge)
	 store.Options.Path = "/"
	 store.Options.HttpOnly = true   // HttpOnly should always be enabled
	 store.Options.Secure = isProd

	 gothic.Store = store


	// Setup authorization
	goth.UseProviders(		
		spotify.New(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_SECRET"), "http://localhost:3000/auth/spotify/callback"),
	)


}
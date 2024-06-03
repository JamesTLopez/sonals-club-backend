package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sonalsguild/helpers"

	"github.com/golang-jwt/jwt/v5"
)

// HTTP handler accessing data from the request context.
func VerifyToken(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		// Parse the jwt from the cookie if exists
		authCookies, err := request.Cookie("auth_token")
		if err!= nil {
            fmt.Println("Error retrieving auth_token:", err)
			helpers.ErrorJson(writer,err)
			return
        }

		tokenString := authCookies.Value

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			log.Fatal(err)
		}

		claims, ok := token.Claims.(jwt.MapClaims);
		
		if ok {
		} else {
			fmt.Println("Error retrieving claims:", err)
			helpers.ErrorJson(writer,err)
		}
		
		ctx := context.WithValue(request.Context(), "user",claims) 
		next.ServeHTTP(writer, request.WithContext(ctx))
    })
}
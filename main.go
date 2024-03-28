package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage (w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, "Welcome to the jungle!");
	fmt.Printf("home page endpoint")

}

func main() {
	log.Println("Hai")
	fmt.Println("FALS")
	http.HandleFunc( "/", homePage)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
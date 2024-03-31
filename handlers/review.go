package handlers

import (
	"fmt"
	"net/http"
)

type Order struct {

}


func (o *Order) CreateReview(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Review created");
}

func (o *Order) ReviewList(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Retrieve Review List");
}

func (o *Order) GetReviewById(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Retrieve Review");
}

func (o *Order) UpdateReviewById(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Update Review");
}

func (o *Order) DeleteReviewById(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Delete Review");
}
package handlers

import (
	"fmt"
	"net/http"
)

type Review struct {

}


func (o *Review) CreateReview(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Review created");
}

func (o *Review) ReviewList(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Retrieve Review List");
}

func (o *Review) GetReviewById(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Retrieve Review");
}

func (o *Review) UpdateReviewById(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Update Review");
}

func (o *Review) DeleteReviewById(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Delete Review");
}
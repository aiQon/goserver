package controllers

import (
	"log"
	"net/http"
)

func UsersRegisterPost(w http.ResponseWriter, r *http.Request) *InternalApiError {
	log.Println("received POST")
	return nil
}

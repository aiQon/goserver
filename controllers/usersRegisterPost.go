package controllers

import (
	"log"
	"net/http"
	"github.com/aiQon/goserver/models"
)

func UsersRegisterPost(w http.ResponseWriter, r *http.Request) *InternalApiError {
	log.Println("received POST")
	return nil
}

package main

import (
	"encoding/json"
	"net/http"
	"github.com/aiQon/goserver/controllers"
	"github.com/aiQon/goserver/models"
	"github.com/gorilla/mux"
)

var statusCodeMessages map[int]string = map[int]string{
	1: "Request message does not match",
}

type ControllerHandler func(http.ResponseWriter, *http.Request) *InternalApiError

func (fn ControllerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		report := &ExternalApiError{StatusCode: e.StatusCode, Message: statusCodeMessages[e.StatusCode]}
		responseJson, _ := json.Marshal(report)
		w.WriteHeader(e.HttpCode)
		w.Write(responseJson)
	}
}

func use(h ControllerHandler, middleware ...func(http.Handler) http.Handler) http.Handler {
	var res http.Handler = h
	for _, m := range middleware {
		res = m(res)
	}
	return res
}

func main() {
	router := mux.NewRouter()

	router.Handle("/users/register", use(UsersRegisterPost)).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/user/{userId}", getUser)

	http.ListenAndServe(":8080", r)
}

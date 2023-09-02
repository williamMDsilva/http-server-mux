package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]
	userName := getUserById(userId)

	if userName == "" {
		fmt.Fprintf(w, "User Not Found")
		return
	}

	fmt.Fprintf(w, "User requested %s", userName)

}

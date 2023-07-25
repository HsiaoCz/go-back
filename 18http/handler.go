package main

import (
	"fmt"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Fprintln(w,username,password)
}

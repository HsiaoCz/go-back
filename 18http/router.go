package main

import "net/http"

func Router() {
	http.HandleFunc("/user", HandleUser)
	http.HandleFunc("/user", HandleGetUser)
}

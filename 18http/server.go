package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	Router()
	srv := http.Server{
		Handler:      nil,
		Addr:         ":9091",
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
	log.Fatal(srv.ListenAndServe())
}

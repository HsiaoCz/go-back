package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// hello

const (
	addr = "127.0.0.1:9911"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		log.Fatal(err)
	}
	username := r.URL.Query().Get("username")
	tmpl.Execute(w, username)
}

func main() {
	http.HandleFunc("/api/hello", HandleHello)
	srv := http.Server{
		Handler:      nil,
		Addr:         addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

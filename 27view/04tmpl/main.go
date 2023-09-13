package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	addr = "127.0.0.1:9991"
)

type H map[string]string

type Book struct {
	BookName string `json:"book_name"`
	Auther   string `json:"auther"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Type     string `json:"type"`
}

func main() {
	http.HandleFunc("/api/book", HandleExecTemplate)
	srv := http.Server{
		Handler:      nil,
		Addr:         addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

func HandleExecTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./book.html")
	if err != nil {
		log.Fatal(err)
	}
	switch r.Method {
	case "POST":
		book := new(Book)
		err = json.NewDecoder(r.Body).Decode(book)
		if err != nil {
			w.Header().Set("Content-Type", "application")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(H{
				"Message": "出错了",
				"Error":   err.Error(),
			})
			return
		}
		tmpl.Execute(w, book)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

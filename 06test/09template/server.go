package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Content  string `json:"content" form:"content"`
}

const addr = "127.0.0.1:9091"

func main() {
	http.HandleFunc("/user", HandleUser)
	http.HandleFunc("/user/post-json", HandleUserPostJson)
	fmt.Println("the server is running on port:", addr)
	http.ListenAndServe(addr, nil)
}
func HandleUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	content := r.FormValue("content")

	user := User{
		Username: username,
		Password: password,
		Content:  content,
	}
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, user)
}

func HandleUserPostJson(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, user)
}

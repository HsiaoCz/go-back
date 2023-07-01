package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Gender string
	Info   map[string]string
}

func main() {
	http.HandleFunc("/user", sayHello)
	http.ListenAndServe(":9090", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := User{
		Name:   "Alex",
		Gender: "nan",
		Age:    11,
		Info: map[string]string{
			"hello":   "hi",
			"bob":     "alic",
			"message": "bil",
		},
	}
	tmpl.Execute(w, user)
}

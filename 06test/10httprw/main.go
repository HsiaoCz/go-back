package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	Addr = "127.0.0.1:9091"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// http request 和 response
// *http.Request 表示请求，可以从中获取请求的参数
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var user User
	// 获取form参数
	username := r.FormValue("username")
	password := r.FormValue("password")
	user = User{
		Username: username,
		Password: password,
	}
	// http.ResponseWriter表示响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func main() {
	http.HandleFunc("/user/register", UserRegister)
	srv := http.Server{
		Handler:      nil,
		Addr:         Addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"log"
	"net/http"
	"time"
)

// 这里使用原生标志库实现一个拿到动态路由上的用户id

const (
	Addr = "127.0.0.1:9091"
)

type User struct {
	Identity int      `json:"identity"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Content  string   `json:"content"`
	Article  []string `json:"article"`
}

type UserRegister struct {
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	RePassword string `json:"re_password" form:"re_password"`
}

func main() {
	http.HandleFunc("/user/register", HandleUserRegister)
	srv := http.Server{
		Handler:      nil,
		Addr:         Addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

func HandleUserRegister(w http.ResponseWriter, r *http.Request) {

}

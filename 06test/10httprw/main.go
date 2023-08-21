package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// 先判断一下是否是get请求
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		userid, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Message:错误的id"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("id:%d", userid)))
	default:
		w.WriteHeader(404)
	}
}

func main() {
	http.HandleFunc("/user/register", UserRegister)
	http.HandleFunc("/user", GetUserById)
	srv := http.Server{
		Handler:      nil,
		Addr:         Addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

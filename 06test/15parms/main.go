package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 这里使用原生标志库实现一个拿到动态路由上的用户id

const (
	Addr = "127.0.0.1:9091"
)

type H map[string]any

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
	http.HandleFunc("/user/getuser", HandleGetUserByID)
	srv := http.Server{
		Handler:      nil,
		Addr:         Addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

func HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// 这种方式是使用form来传递参数
		// username := r.FormValue("username")
		// password := r.FormValue("password")
		// re_password := r.FormValue("re_password")
		var userR UserRegister
		json.NewDecoder(r.Body).Decode(&userR)
		if userR.Password != userR.RePassword {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(H{
				"message": "password和re_password不相等,请重新输入",
			})
			return
		}
		user := User{
			Identity: GenIdnetity(),
			Username: userR.Username,
			Password: userR.Password,
			Content:  "Hello My man",
			Article:  []string{"遮天", "武动乾坤", "斗破苍穹"},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&user)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// 这里使用这种方式实现是错误的
// 因为要匹配动态路由
// 需要对路由进行匹配
// 这里我们需要自己实现一个多路复用器
// 或者说自己的handler

func HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path
		pathStrings := strings.Split(path, "/")
		id := pathStrings[len(pathStrings)-1]
		userId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		user := User{
			Identity: userId,
			Username: "andy",
			Password: "不告诉你",
			Content:  "Hello My Man",
			Article:  []string{"仙逆", "校花还在贴身高手"},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&user)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func GenIdnetity() int {
	randen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randen.Intn(1000000)
}

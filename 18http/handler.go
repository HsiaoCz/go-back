package main

import (
	"fmt"
	"net/http"
)

// formValue可以拿到form数据
func HandleUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Fprintln(w, username, password)
}

// 通过r.URL.Query.Get() 拿到Url上的请求参数
func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	w.Write([]byte(id))
}

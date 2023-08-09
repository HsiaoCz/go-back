package router

import (
	handl "go-back/18http/handler"
	"net/http"
)

const addr = "127.0.0.1:9091"

func RegisterRouter() {
	http.HandleFunc("/user", handl.HandleUser)
	http.HandleFunc("/user/id", handl.HandleGetUser)
	http.ListenAndServe(addr, nil)
}

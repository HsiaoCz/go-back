package main

import (
	"go-back/26mansys/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Addr = "127.0.0.1:9093"
)

// 来做一个毕设一样的东西
// 图书馆里系统
func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.RegisterRouter(r)
	srv := http.Server{
		Handler:      r,
		Addr:         Addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

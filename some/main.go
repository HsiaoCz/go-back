package main

import (
	"go-back/some/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Addr = "127.0.0.1:9091"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.RegisterRouter(r)
	srv := http.Server{
		Handler:      r,
		Addr:         Addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

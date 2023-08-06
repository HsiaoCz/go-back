package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", GetUser)
	r.Run(":9090")
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "hello",
	})
}

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/user/register", HandleUserRegister)

}
func HandleUserRegister(c *gin.Context) {

}

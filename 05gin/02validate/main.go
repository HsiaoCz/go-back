package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin中的参数校验
// 这种校验方式有一个问题
// 1、返回的错误是英文的
// 2、会将结构体的字段信息返回回去

// 所以这里需要解决这两个问题
// 1、注册一个翻译器
// 2、对validate返回的错误信息进行处理，去掉敏感信息再返回给前端
func main() {
	r := gin.Default()
	r.POST("/user", ValidateUser)
	r.Run("127.0.0.1:9091")
}

type User struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required,email"`
}

func ValidateUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

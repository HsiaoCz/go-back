package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserRegister(c *gin.Context) {
	var userR UserR
	err := c.BindJSON(userR)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if userR.Password != userR.RePassword {
		c.JSON(http.StatusOK, gin.H{
			"Message": "请检查确认密码",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "注册成功",
	})
}

func HandleUserLogin(c *gin.Context) {
	var userl UserL
	err := c.BindJSON(userl)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "登录成功",
	})
}

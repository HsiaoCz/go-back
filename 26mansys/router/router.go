package router

import (
	"go-back/26mansys/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/user/register", controller.HandleUserRegister)
	r.POST("/user/login")
}

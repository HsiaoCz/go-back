package router

import (
	"go-back/some/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/api/v1/book/:book_id", controller.GetBookBuyID)
}

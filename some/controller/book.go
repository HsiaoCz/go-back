package controller

import (
	"go-back/some/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookBuyID(c *gin.Context) {
	id := c.Param("book-id")
	book_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "输入的信息有误",
		})
		return
	}
	book := model.Book{
		BookIdentity: int64(book_id),
		BookName:     "老衲法号秃驴",
		Auther:       "andy",
		Title:        "那些不得不说的故事",
		Content:      "hello my man",
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "成功找到信息",
		"Data":    book,
	})
}

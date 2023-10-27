package controller

import (
	"go-back/30fiber/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBookBuyID(c *fiber.Ctx) error {
	id := c.Params("book-id")
	book_id, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "输入的信息有误",
		})
		return err
	}
	book := model.Book{
		BookIdentity: int64(book_id),
		BookName:     "老衲法号秃驴",
		Auther:       "andy",
		Title:        "那些不得不说的故事",
		Content:      "hello my man",
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "成功找到信息",
		"Data":    book,
	})
}

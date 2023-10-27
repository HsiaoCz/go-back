package router

import (
	"go-back/30fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func RegRouter(r *fiber.App) {
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			book := v1.Group("/book")
			{
				book.Get("/:book-id", controller.GetBookBuyID)
			}
		}
	}
}

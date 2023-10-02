package main

import (
	"go-back/30fiber/router"

	"github.com/gofiber/fiber/v2"
)

const (
	addr = "127.0.0.1:3301"
)

func main() {
	app := fiber.New()
	router.RegRouter(app)
	app.Listen(addr)
}

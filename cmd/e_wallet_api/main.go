package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func main()  {
	fmt.Println("ISHLADI!")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	fmt.Errorf("", app.Listen(":8001"))
}

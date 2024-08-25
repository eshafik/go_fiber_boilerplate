package main

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber_boilerplate/config"
)

func main() {
	app := fiber.New()
	config.AppSettings()
	AddRoutes(app)
	app.Listen(":3000")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

package main

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber_boilerplate/Apps/User"
)

func AddRoutes(app *fiber.App) {
	// Add app wise prefix and app routers
	userAppRoute := app.Group("/user")
	User.AddAppRoutes(userAppRoute)
}

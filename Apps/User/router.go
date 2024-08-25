package User

import "github.com/gofiber/fiber/v2"

func AddAppRoutes(router fiber.Router) {
	router.Post("/token", getUserToken)
}

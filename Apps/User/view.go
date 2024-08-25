package User

import "github.com/gofiber/fiber/v2"

func getUserToken(cx *fiber.Ctx) error {
	return cx.Status(200).JSON(fiber.Map{"success": true})
}

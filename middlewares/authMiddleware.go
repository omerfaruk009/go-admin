package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go-admin2/util"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"mesaj": "unauthorized",
		})
	}
	return c.Next()
}

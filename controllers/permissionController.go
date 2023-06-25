package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin2/database"
	"go-admin2/models"
)

func AllPermissions(c *fiber.Ctx) error {

	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)

}

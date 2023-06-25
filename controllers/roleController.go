package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin2/database"
	"go-admin2/models"
	"strconv"
)

func AllRole(c *fiber.Ctx) error {

	var roles []models.Role

	database.DB.Order("id ASC").Find(&roles)

	return c.JSON(roles)

}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)
	return c.JSON(role)
}

type temp struct {
	Name        string `json:"name"`
	Permissions []uint `json:"permissions"`
}

func UpdateRole(c *fiber.Ctx) error {
	var roleDto temp

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	var dbmodel models.Role
	if err := database.DB.Model(&dbmodel).Where("id = ?", id).First(&dbmodel).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&dbmodel).Association("Permissions").Clear(); err != nil {
		return err
	}

	for _, permission := range roleDto.Permissions {
		dbmodel.Permissions = append(dbmodel.Permissions, models.Permission{Id: permission})
	}

	if err := database.DB.Save(&dbmodel).Error; err != nil {
		return err
	}

	return c.JSON(dbmodel)

}

func DeleteRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}

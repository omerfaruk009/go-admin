package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-admin2/database"
	"go-admin2/routs"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8080",
		AllowMethods:     "GET,PUT,POST,DELETE,PATCH,OPTIONS",
	}))
	database.Connect()
	routs.Setupp(app)
	log.Fatal(app.Listen(":8000"))
}

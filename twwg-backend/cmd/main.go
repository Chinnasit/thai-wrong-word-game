package main

import (
	"Chinnasit/pkg/common/db"
	"Chinnasit/pkg/questions"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	dbHandler := db.Init()
	questions.RegisterRoutes(app, dbHandler)

	app.Listen(":8000")
}

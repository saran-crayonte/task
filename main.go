package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/saran-crayonte/task/database"
	_ "github.com/saran-crayonte/task/docs"
	"github.com/saran-crayonte/task/routes"
)

//	@title			Task Management API
//	@Description	This is a sample API for managing tasks.
//	@version		1.0
//	@contact.name	Saran
//	@contact.url	github.com/saran-crayonte/
//	@contact.email	saran.kumaresan@crayonte.com
//	@host			localhost:8080
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	database.ConnectDB()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}

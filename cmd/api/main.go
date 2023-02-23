package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/victorbischoff/structs-api/pkg/database"
	"github.com/victorbischoff/structs-api/pkg/routers"
)

func init() {
	database.InitDb()

}

// Create db connection...
func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "API v1",
		AppName:      "User Crud App v1.0.1",
	})

	routers.ApiRouter(app)
	log.Fatal(app.Listen(":1312"))
}

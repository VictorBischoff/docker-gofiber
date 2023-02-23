package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorbischoff/structs-api/pkg/handlers"
)

func ApiRouter(app *fiber.App) {
	app.Route("/api", func(router fiber.Router) {

		v1 := router.Group("/v1")
		v1.Get("/healthcheck", func(c *fiber.Ctx) error {
			return c.SendString("We're good!")
		}).Name("healthcheck")

		users := v1.Group("/users")
		users.Get("/", handlers.GetAllUsers)
		users.Post("/", handlers.NewUser)
	},"api")
}

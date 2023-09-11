package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/ramdeoangh/goswagger/controllers"
	"github.com/ramdeoangh/goswagger/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8000/swagger/oauth2-redirect.html",
	}))

	app.Mount("/api", micro)
	app.Use(logger.New())
	micro.Use(cors.New())

	micro.Route("/notes", func(router fiber.Router) {
		router.Post("/", controllers.CreateNoteHandler)
		router.Get("", controllers.FindNotes)
		router.Get("/search", controllers.FindNotesByDate)
	})
	micro.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", controllers.DeleteNote)
		router.Get("", controllers.FindNoteById)
		router.Patch("", controllers.UpdateNote)
	})
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, MySQL, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}

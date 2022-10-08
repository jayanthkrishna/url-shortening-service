package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jayanthkrishna/url-shortening-service/routes"
	"github.com/joho/godotenv"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1/", routes.ShortenURL)
}
func main() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	SetupRoutes(app)

	app.Listen(os.Getenv("APP_PORT"))

}

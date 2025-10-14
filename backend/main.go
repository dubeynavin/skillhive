package main

import (
	"log"
	

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/yourusername/skillhiveai-backend/backend/config"
	"github.com/yourusername/skillhiveai-backend/backend/routes"
)

func main() {
	_ = godotenv.Load()
	config.Connect()

	app := fiber.New(fiber.Config{})

	// CORS for frontend
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(200)
		}
		return c.Next()
	})

	// root health endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SkillHiveAI backend is running!")
	})

	routes.Setup(app)

	log.Println("Listening on :8080")
	log.Fatal(app.Listen(":8080"))
}

package routes

import (
	"github.com/yourusername/skillhiveai-backend/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/auth/register", controllers.Register)
	api.Post("/auth/login", controllers.Login)
	api.Post("/auth/forgot", controllers.ForgotPassword)
	api.Post("/auth/reset", controllers.ResetPassword)

	api.Post("/tasks", controllers.CreateTask)
	api.Get("/tasks", controllers.GetTasks)
}

package controllers

import (
	"context"
	"time"
	"github.com/yourusername/skillhiveai-backend/backend/config"
	"github.com/yourusername/skillhiveai-backend/backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTask(c *fiber.Ctx) error {
	var t models.Task
	if err := c.BodyParser(&t); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	t.CreatedAt = time.Now()
	collection := config.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, t)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "db error"})
	}
	return c.JSON(fiber.Map{"message": "task posted"})
}

func GetTasks(c *fiber.Ctx) error {
	collection := config.DB.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "db error"})
	}
	var tasks []models.Task
	if err := cur.All(ctx, &tasks); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "db error"})
	}
	return c.JSON(tasks)
}

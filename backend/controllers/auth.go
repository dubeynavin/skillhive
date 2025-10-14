package controllers

import (
	"context"
	"time"
	"os"
	"github.com/yourusername/skillhiveai-backend/backend/config"
	"github.com/yourusername/skillhiveai-backend/backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Register
func Register(c *fiber.Ctx) error {
	type req struct { Name, Email, Password string }
	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	// check exist
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, _ := collection.CountDocuments(ctx, bson.M{"email": body.Email})
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "email already registered"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user := models.User{
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "db error"})
	}
	return c.JSON(fiber.Map{"message": "registered"})
}

// Login
func Login(c *fiber.Ctx) error {
	type req struct { Email, Password string }
	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var u models.User
	if err := collection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&u); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// create JWT
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{"email": u.Email, "name": u.Name, "exp": time.Now().Add(time.Hour * 72).Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "token error"})
	}
	return c.JSON(fiber.Map{"token": t, "user": u})
}

// Forgot password (generates a reset token and returns a link - in production send email)
func ForgotPassword(c *fiber.Ctx) error {
	type req struct { Email string }
	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var u models.User
	if err := collection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&u); err != nil {
		// do not reveal whether email exists
		return c.JSON(fiber.Map{"message": "If email exists, reset link sent"})
	}
	// create token
	token := uuid.New().String()
	exp := time.Now().Add(30 * time.Minute)
	_, err := collection.UpdateOne(ctx, bson.M{"email": body.Email}, bson.M{"$set": bson.M{"reset_token": token, "reset_expires_at": exp}})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "db error"})
	}
	// Build reset URL and return it (for demo). In prod, email this to the user.
	resetURL := os.Getenv("FRONTEND_URL") + "/reset-password?token=" + token + "&email=" + body.Email
	return c.JSON(fiber.Map{"message": "reset link generated", "reset_url": resetURL})
}

// Reset password
func ResetPassword(c *fiber.Ctx) error {
	type req struct { Email, Token, NewPassword string }
	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var u models.User
	if err := collection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid token or email"})
	}
	if u.ResetToken == "" || u.ResetToken != body.Token || time.Now().After(u.ResetExpiresAt) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid or expired token"})
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	_, err := collection.UpdateOne(ctx, bson.M{"email": body.Email}, bson.M{"$set": bson.M{"password_hash": string(hash)}, "$unset": bson.M{"reset_token": "", "reset_expires_at": ""}})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "db error"})
	}
	return c.JSON(fiber.Map{"message": "password reset successful"})
}

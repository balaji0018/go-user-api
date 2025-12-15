package main

import (
	"context"
	"go-user-api/internal/handler"
	"go-user-api/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. Connect to Database
	dbUrl := "postgres://postgres:system@localhost:5432/userdb?sslmode=disable"

	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	defer conn.Close()

	// 2. Initialize Layers (Service -> Handler)
	userService := service.NewUserService(conn)
	userHandler := handler.NewUserHandler(userService)

	// 3. Setup Fiber App (The Web Server)
	app := fiber.New()

	// 4. Define Routes
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUser)
	app.Get("/users", userHandler.GetAllUsers)
	app.Delete("/users/:id", userHandler.DeleteUser)

	// 5. Start Server on Port 3000
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Welcome to Go User API! Use /users endpoints to interact with the API.")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("⚠️ Route not found. Please use /users endpoints.")
	})

	log.Println("Server is running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}

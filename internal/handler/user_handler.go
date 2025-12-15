package handler

import (
	"go-user-api/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserHandler struct {
	service *service.UserService
}

// NewUserHandler creates the handler
func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// createUser handles Post/users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	//Read json bodysent by the user
	type Request struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	//Call service
	user, err := h.service.CreateUser(c.Context(), req.Name, req.Dob)
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	//send back created user
	return c.Status(201).JSON(user)
}

// Getuser handles get /user/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	//get the id from the url
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	//call the service
	user, err := h.service.GetUser(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	//send back the user(with age)
	return c.JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.JSON(users)
}
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = h.service.DeleteUser(c.Context(), int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendStatus(204) // No Content
}

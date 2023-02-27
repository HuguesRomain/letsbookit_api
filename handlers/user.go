package handlers

import (
	"github.com/HuguesRomain/letsbookit_api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userHandler struct {
	repo models.UserRepository
}

func NewUserHandler(db *gorm.DB) UserHandler {
	repo := models.NewUserRepository(db)
	return &userHandler{repo}
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	user := &models.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   req.Password,
		IsMerchant: req.IsMerchant,
	}
	err := h.repo.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	user, err := h.repo.FindByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid username or password"})
	}

	if user.Password != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid username or password"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged in successfully"})
}

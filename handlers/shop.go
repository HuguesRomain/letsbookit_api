package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/HuguesRomain/letsbookit_api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type ShopHandler interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type shopHandler struct {
	repo models.ShopRepository
	db   *gorm.DB
}

func NewShopHandler(db *gorm.DB) ShopHandler {
	repo := models.NewShopRepository(db)
	return &shopHandler{repo, db}
}

func (h *shopHandler) Create(c *fiber.Ctx) error {
	fmt.Println(h.getUserFromContext(c))
	user, err := h.getUserFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	if !user.IsMerchant {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "User not authorized to create shop"})
	}

	var req models.CreateShopRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	shop := &models.Shop{
		Name:    req.Name,
		Address: req.Address,
		UserID:  user.ID,
	}
	if err := h.repo.CreateShop(shop); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create shop"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Shop created successfully"})
}

func (h *shopHandler) Get(c *fiber.Ctx) error {
	shopID := c.Params("id")

	shop, err := h.repo.Get(shopID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Shop not found"})
	}

	return c.Status(fiber.StatusOK).JSON(shop)
}

func (h *shopHandler) GetAll(c *fiber.Ctx) error {
	shops, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get shops"})
	}

	return c.Status(fiber.StatusOK).JSON(shops)
}

func (h *shopHandler) getUserFromContext(c *fiber.Ctx) (*models.User, error) {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return nil, errors.New("authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid authorization header")
	}
	if parts[0] != "Bearer" {
		return nil, errors.New("invalid authorization header")
	}

	token := parts[1]
	user, err := models.GetUserFromToken(token, h.db)
	if err != nil {
		return nil, err
	}

	return user, nil
}

package handlers

import (
	"strconv"

	"github.com/HuguesRomain/letsbookit_api/models"
	"github.com/gofiber/fiber/v2"
)

func (h *serviceHandler) Create(c *fiber.Ctx) error {
	shopID := c.Params("shopID")

	shopIDUint, err := strconv.ParseUint(shopID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid shop ID"})
	}

	shop := &models.Shop{}
	if err := h.db.First(&shop, shopIDUint).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Shop not found"})
	}

	var req models.CreateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	service := &models.Service{
		Name:      req.Name,
		Price:     req.Price,
		Duration:  req.Duration,
		Available: true,
		ShopID:    shop.ID,
	}

	if err := h.db.Create(&service).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create service"})
	}

	if err := h.db.Model(&service).Preload("Shop").First(&service).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to load shop"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Service created successfully", "service": service})
}

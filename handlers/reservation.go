package handlers

import (
	"github.com/HuguesRomain/letsbookit_api/models"
	"github.com/gofiber/fiber/v2"
)

type ReservationHandler interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type reservationHandler struct {
	repo models.ReservationRepository
}

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Reservation struct {
	ID        uint      `gorm:"primary_key"`
	Date      time.Time `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignkey:UserID"`
	ShopID    uint      `gorm:"not null"`
	Shop      Shop      `gorm:"foreignkey:ShopID"`
	ServiceID uint      `gorm:"not null"`
	Service   Service   `gorm:"foreignkey:ServiceID"`
}

type ReservationRepository interface {
	CreateReservation(shop *Shop) error
	Get(id string) (*Shop, error)
	GetAll() ([]Shop, error)
}

type reservationRepo struct {
	db *gorm.DB
}

type CreateReservationRequest struct {
	Date   time.Time `gorm:"type:date"`
	ShopID uint      `gorm:"not null"`
}

package models

import (
	"time"
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

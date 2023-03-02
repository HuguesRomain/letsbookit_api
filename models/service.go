package models

type Service struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"not null"`
	Price        int    `gorm:"not null"`
	ShopID       uint   `gorm:"not null"`
	Shop         Shop   `gorm:"foreignkey:ShopID"`
	Reservations []Reservation
}

package models

import (
	"github.com/jinzhu/gorm"
)

type Shop struct {
	gorm.Model
	Name         string        `gorm:"not null"`
	Address      string        `gorm:"not null"`
	UserID       uint          `gorm:"not null"`
	User         *User         `gorm:"foreignkey:UserID"`
	Reservations []Reservation `gorm:"ForeignKey:ShopID"`
	Services     []Service     `gorm:"ForeignKey:ShopID"`
}

type ShopRepository interface {
	CreateShop(shop *Shop) error
	Get(id string) (*Shop, error)
	GetAll() ([]Shop, error)
}

type shopRepo struct {
	db *gorm.DB
}

type CreateShopRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepo{db}
}

func (repo *shopRepo) CreateShop(shop *Shop) error {
	return repo.db.Create(shop).Error
}

func (repo *shopRepo) Get(id string) (*Shop, error) {
	var shop Shop
	err := repo.db.Preload("User").Preload("Reservations").Preload("Services").Where("id = ?", id).First(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (repo *shopRepo) GetAll() ([]Shop, error) {
	var shops []Shop
	err := repo.db.Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return shops, nil
}

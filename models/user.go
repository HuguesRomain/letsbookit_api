package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	IsMerchant bool   `gorm:"not null"`
	UUID       string `gorm:"unique;not null"`
}

type RegisterRequest struct {
    Username   string `json:"username"`
    Email      string `json:"email"`
    Password   string `json:"password"`
    IsMerchant bool   `json:"is_merchant"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}


type UserRepository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (repo *userRepo) Create(user *User) error {
	user.UUID = uuid.New().String()
	return repo.db.Create(user).Error
}

func (repo *userRepo) FindByUsername(username string) (*User, error) {
	var user User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

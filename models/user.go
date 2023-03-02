package models

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	IsMerchant   bool   `gorm:"not null"`
	Shops        []Shop
	Services     []Service
	Reservations []Reservation
	UUID         string `gorm:"unique;not null"`
}

type RegisterRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsMerchant bool   `json:"isMerchant"`
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

func FindUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserFromToken(tokenString string, db *gorm.DB) (*User, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get the secret key from the environment variables
		secretKey := []byte("my-secret-key")

		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Get the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	userID, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid user ID in token claims")
	}

	// Get the user from the database
	user, err := FindUserByID(db, uint(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}

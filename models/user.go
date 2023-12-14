package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

const (
	Admin   = "Admin"
	General = "General"
)

// User represents the user model
type User struct {
	ID        int            `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone"`
	Username  string         `json:"username"`
	Role      string         `json:"role"`
	IsActive  bool           `json:"is_active"`
	Tasks     []Task         `json:"tasks"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt gorm.DeletedAt `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// JWTCustomClaims represents the custom claims for JWT
type JWTCustomClaims struct {
	UserID        int    `json:"user_id"`
	UserFirstName string `json:"user_first_name"`
	UserLastName  string `json:"user_last_name"`
	jwt.StandardClaims
}

// LoginRequest represents the structure of the login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

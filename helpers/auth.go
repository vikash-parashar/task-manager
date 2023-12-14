// helpers/auth_helpers.go

package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/vikash-parashar/task-remainder/models"
)

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash checks whether the provided password matches the hashed password
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return err == nil
}

// GenerateJWTToken generates a JWT token for the given user
func GenerateJWTToken(user *models.User) (string, error) {
	// Create a new JWT claims
	claims := &models.JWTCustomClaims{
		UserID:        user.ID,
		UserFirstName: user.FirstName,
		UserLastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expires in 7 days
		},
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	// Sign the token with a secret key
	signedToken, err := token.SignedString(secretKey) // Replace with your actual secret key
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// GetUsernameFromEmail returns the username part of an email address
func GetUsernameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return email
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Add your JWT secret key here for validation
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

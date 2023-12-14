package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/vikash-parashar/task-remainder/controllers"
// 	"github.com/vikash-parashar/task-remainder/helpers"
// 	"github.com/vikash-parashar/task-remainder/models"
// )

// // GetUserIDFromContext()

// // Import necessary packages and modules

// // RegisterUser creates a new user based on the registration request
// func RegisterUser(w http.ResponseWriter, r *http.Request) {
// 	var user models.User

// 	// Decode the JSON request body into the user struct
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
// 		return
// 	}

// 	// Trim the text after @ in the email address
// 	user.Email = strings.Split(user.Email, "@")[0]

// 	// Set the user's role to "General" for simplicity
// 	user.Role = models.General

// 	// Hash the user's password before storing it in the database
// 	hashedPassword, err := helpers.HashPassword(user.Password)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	user.Password = hashedPassword

// 	// Create the user in the database
// 	err = controllers.CreateUser(&user)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }

// // LoginUser handles the login request and generates a JWT token on successful login
// func LoginUser(w http.ResponseWriter, r *http.Request) {
// 	var loginRequest models.LoginRequest

// 	// Decode the JSON request body into the loginRequest struct
// 	err := json.NewDecoder(r.Body).Decode(&loginRequest)
// 	if err != nil {
// 		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
// 		return
// 	}

// 	// Try to find the user by username or email
// 	user, err := controllers.GetUserByUsernameOrEmail(loginRequest.Username)
// 	if err != nil {
// 		http.Error(w, "No user found with the given username or email", http.StatusNotFound)
// 		return
// 	}

// 	// Compare the provided password with the stored hashed password
// 	if !helpers.CheckPasswordHash(loginRequest.Password, user.Password) {
// 		http.Error(w, "Incorrect password. Please try again", http.StatusUnauthorized)
// 		return
// 	}

// 	// Generate a JWT token
// 	token, err := helpers.GenerateJWTToken(user)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the token as a cookie in the response
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "token",
// 		Value:    token,
// 		HttpOnly: true,
// 		SameSite: http.SameSiteStrictMode,
// 		Secure:   true,                               // Set to true if using HTTPS
// 		Expires:  time.Now().Add(time.Hour * 24 * 7), // Token expires in 7 days
// 	})

// 	w.WriteHeader(http.StatusOK)
// }

// // GetCurrentUser retrieves the current user based on the JWT token
// func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
// 	tokenString := r.Header.Get("Authorization")
// 	if tokenString == "" {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	// Parse and validate the JWT token
// 	token, err := helpers.ParseToken(tokenString)
// 	if err != nil {
// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Extract claims from the token
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
// 		return
// 	}

// 	// Retrieve user ID from claims
// 	userID, ok := claims["user_id"].(float64)
// 	if !ok {
// 		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Convert user ID to int
// 	userIDInt := int(userID)

// 	// Retrieve user from the database based on the user ID
// 	user, err := controllers.GetUserByID(userIDInt)
// 	if err != nil {
// 		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
// 		return
// 	}

// 	// Respond with the user details
// 	helpers.RespondJSON(w, user)
// }

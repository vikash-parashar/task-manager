// config/config.go
package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vikash-parashar/task-remainder/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Load environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Initialize database connection
func InitDB() {
	LoadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbSSLMode)

	var err error

	// Set up GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Enable color
		},
	)

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Automatically migrate the schema, creating tables
	err = DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal(err)
	}
}

// Create the tasks table
// func createTasksTable() {
// 	err := DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS tasks (
// 			id SERIAL PRIMARY KEY,
// 			user_id INT,
// 			title VARCHAR(255),
// 			description TEXT,
// 			priority INT,
// 			due_date TIMESTAMP,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			FOREIGN KEY (user_id) REFERENCES users(id)
// 		)
// 	`).Error

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// createUserTable creates the user table in the database
// func createUserTable() {
// 	err := DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS users (
// 			id SERIAL PRIMARY KEY,
// 			first_name VARCHAR(255),
// 			last_name VARCHAR(255),
// 			email VARCHAR(255),
// 			password VARCHAR(255),
// 			phone VARCHAR(15),
// 			username VARCHAR(255),
// 			role VARCHAR(50),
// 			is_active BOOLEAN,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		)
// 	`).Error

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// Close the database connection
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.Close()
	}
}

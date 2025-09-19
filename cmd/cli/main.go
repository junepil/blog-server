package main

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 2. Define and parse command-line flags
	email := flag.String("email", "", "Admin user's email")
	password := flag.String("password", "", "Admin user's password")
	flag.Parse()

	if *email == "" || *password == "" {
		fmt.Println("Both --email and --password flags are required.")
		os.Exit(1)
	}

	// 3. Connect to the database
	database.Connect()
	db := database.DB

	// 4. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// 5. Create the admin user
	adminUser := models.User{
		Username: "admin", // Default username for the first admin
		Email:    *email,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	fmt.Println("Admin user created successfully!")
}
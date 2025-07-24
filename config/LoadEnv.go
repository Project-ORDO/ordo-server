package config

import (
	"log"

	"github.com/joho/godotenv"
)

/* func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env fil")
		return
	}
} */


func LoadEnvFile() {
	// Only load .env if not running on Render
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("⚠️  Warning: .env file not found, skipping...")
		} else {
			log.Println("✅ .env file loaded successfully.")
		}
	} else {
		log.Println("🌐 Running on Render, skipping .env load")
	}
}
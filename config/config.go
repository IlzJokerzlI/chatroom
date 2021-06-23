package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPort() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(`Error loading .env file`)
	}
	port := os.Getenv(`PORT`)
	return port
}

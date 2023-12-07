package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		if strings.EqualFold(os.Getenv("APP_DEBUG"), "true") {
			NewLogger().Fatalln("config file not loaded")
		}
	}
}
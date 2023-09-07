package configs

import (
	"log"
	"os"
	"server/models"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var Config models.ConfigType

func SetupConfig() {
	godotenv.Load()

	Config = models.ConfigType{
		Port:            os.Getenv("PORT"),
		Cors:            os.Getenv("CORS"),
		AdminPassword:   os.Getenv("ADMIN_PASS"),
		OutlookUser:     os.Getenv("OUTLOOK_USER"),
		OutlookPassword: os.Getenv("OUTLOOK_PASS"),
	}

	if Config.Port == "" {
		Config.Port = "4000"
	}

	if err := validator.New().Struct(&Config); err != nil {
		log.Fatal(err)
	}
}

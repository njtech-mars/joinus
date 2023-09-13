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
	if _, err := os.Stat("data/template.html"); err != nil {
		log.Fatal("Please add data/template.html")
	}

	godotenv.Load()

	Config = models.ConfigType{
		Port:          os.Getenv("PORT"),
		Cors:          os.Getenv("CORS"),
		AdminPassword: os.Getenv("ADMIN_PASS"),
		SMTPUser:      os.Getenv("SMTP_USER"),
		SMTPPass:      os.Getenv("SMTP_PASS"),
		SMTPHOST:      os.Getenv("SMTP_HOST"),
		SMTPPort:      os.Getenv("SMTP_PORT"),
		SMTPTLS:       os.Getenv("SMTP_TLS"),
	}

	if Config.Port == "" {
		Config.Port = "5000"
	}

	if Config.SMTPTLS == "" {
		Config.SMTPTLS = "true"
	}

	if err := validator.New().Struct(&Config); err != nil {
		log.Fatal(err)
	}
}

package models

type ConfigType struct {
	Port          string `validate:"required,number"`
	Cors          string `validate:"omitempty"`
	AdminPassword string `validate:"required"`
	SMTPUser      string `validate:"required,email"`
	SMTPPass      string `validate:"required"`
	SMTPHOST      string `validate:"required"`
	SMTPPort      string `validate:"required,number"`
}

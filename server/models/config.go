package models

type ConfigType struct {
	Port            string `validate:"required,number"`
	Cors            string `validate:"omitempty"`
	AdminPassword   string `validate:"required"`
	OutlookUser     string `validate:"required,email"`
	OutlookPassword string `validate:"required"`
}

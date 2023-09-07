package lib

import (
	"server/configs"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeUser(c *fiber.Ctx, key string) (bool, error) {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Config.AdminPassword), nil
	})
	if err != nil {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}
	// check whether has token expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if expirationTime.Before(time.Now()) {
			return false, keyauth.ErrMissingOrMalformedAPIKey
		}
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

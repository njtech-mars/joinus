package middlewares

import (
	"server/lib"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var Auth = keyauth.New(keyauth.Config{
	KeyLookup: "cookie:joinus_token",
	Validator: lib.AuthorizeUser,
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(401).JSON(models.Response{Success: false, Message: "无效的token"})
	},
})

package middlewares

import (
	"server/lib"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

func Limitter(max int, expr time.Duration) func(*fiber.Ctx) error {
	return skip.New(
		limiter.New(limiter.Config{
			Max:               max,
			Expiration:        expr * time.Second,
			LimiterMiddleware: limiter.SlidingWindow{},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(429).JSON(models.Response{Success: false, Message: "你的操作太频繁啦，请稍后再试"})
			},
		}),
		func(c *fiber.Ctx) bool {
			success, _ := lib.AuthorizeUser(c, c.Cookies("joinus_token"))
			return success
		})
}

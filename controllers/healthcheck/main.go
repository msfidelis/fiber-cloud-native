package healthcheck

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func Get(ctx *fiber.Ctx) {
	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
	})
}

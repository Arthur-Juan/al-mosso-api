package middlwares

import (
	token2 "al-mosso-api/pkg/token"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CheckAuth(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
			"msg": "Forbidden",
		})
	}
	// Remove the "Bearer " prefix from the token
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := token2.CheckToken(tokenString)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": err,
		})
	}
	ctx.Locals("userId", token.ID)
	return ctx.Next()
}

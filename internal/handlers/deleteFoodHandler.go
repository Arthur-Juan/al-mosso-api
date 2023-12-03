package handlers

import (
	"al-mosso-api/config"
	"al-mosso-api/pkg/database/schemas"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

func DeleteFoodHandler(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")
	id, _ := strconv.ParseInt(idStr, 10, 8)

	db := config.GetDb()

	var food schemas.Food
	if err := db.Where("id = ?", id).Find(&food).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NotFound(ctx, "Food not found")
		}
		return InternalServerError(ctx, err)
	}

	db.Delete(&food)
	return NoContent(ctx)

}

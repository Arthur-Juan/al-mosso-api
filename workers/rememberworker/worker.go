package rememberworker

import (
	"al-mosso-api/config"
	"al-mosso-api/pkg/database/schemas"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func Run() {
	db := config.GetDb() // Obtain the database connection outside the loop
	log.Warn(db)
	ticker := time.NewTicker(5 * time.Second)

	for t := range ticker.C {
		dia := time.Now().Day()
		mes := time.Now().Month()
		ano := time.Now().Year()

		current := fmt.Sprintf("%v-%d-%v", ano, mes, dia)
		var appointents []schemas.Appointment
		if err := db.Where("date like ? and verified = true", fmt.Sprintf("%%%v%%", current)).Find(&appointents).Error; err != nil {
			log.Error(err)
		}

		log.Info("appointments", t)
	}
}

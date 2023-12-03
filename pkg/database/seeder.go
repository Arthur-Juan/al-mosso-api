package database

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/pkg/database/schemas"
	logger2 "al-mosso-api/pkg/logger"
	"fmt"
	"gorm.io/gorm"
	"os"
)

var (
	logger = logger2.GetLogger("seeder")
)

func SeedFood() {

	shakshuka, err := entity.NewFood("Shakshuka", 16.90, "                  Prato tradicional do Oriente Médio e do Norte da África, feito com ovos pochê em um molho de tomate temperado com especiarias como pimenta, cominho e páprica, geralmente acompanhado de pimentões e cebolas.")
	if err != nil {
		panic(err)
	}

	babaganush, err := entity.NewFood("babaganush", 20.99, "Prato à base de berinjela assada, misturada com tahine, alho, limão e azeite de oliva, resultando em uma pasta cremosa e saborosa, comum na culinária do Oriente Médio.")
	if err != nil {
		panic(err)
	}

	mujadara, err := entity.NewFood("Mujadara", 30.99, "Prato tradicional do Oriente Médio feito com arroz, lentilhas e cebolas caramelizadas, geralmente temperado com cominho e servido como acompanhamento ou prato principal.")
	if err != nil {
		panic(err)
	}
	tabule, err := entity.NewFood("Tabule", 14.00, "Salada refrescante do Oriente Médio, feita com trigo bulgur, tomates, pepino, salsa, hortelã, cebola e temperada com azeite de oliva e suco de limão, oferecendo um sabor leve e vibrante.")
	if err != nil {
		panic(err)
	}

	muhamara, err := entity.NewFood("Muhamara", 20.00, "Pasta de pimentão vermelho, nozes, alho, azeite de oliva e especiarias, comumente encontrado na culinária do Oriente Médio. Tem uma textura cremosa e sabor picante e agridoce.")
	if err != nil {
		panic(err)
	}

	fatuche, err := entity.NewFood("Fatuche", 25.00, "Salada do Oriente Médio composta de alface, pepino, tomate, rabanete, cebola e pão árabe torrado, temperada com azeite de oliva e sumagre, proporcionando um sabor fresco e crocante.")
	if err != nil {
		panic(err)
	}

	foods := []schemas.Food{
		schemas.Food{
			gorm.Model{},
			*shakshuka,
		},
		schemas.Food{
			gorm.Model{},
			*babaganush,
		},
		schemas.Food{
			gorm.Model{},
			*mujadara,
		},
		schemas.Food{
			gorm.Model{},
			*tabule,
		},
		schemas.Food{
			gorm.Model{},
			*muhamara,
		},
		schemas.Food{
			gorm.Model{},
			*fatuche,
		},
	}

	files, err := os.ReadDir("uploads")
	db := config.GetDb()
	i := 0
	for i < 6 {

		foods[i].ProfilePic = fmt.Sprintf("uploads/%s", files[i+3].Name())
		db.Save(&foods[i])
		logger.Debug(foods[i])
		i += 1
	}
}

func SeedChef() {
	files, err := os.ReadDir("uploads")
	if err != nil {
		panic(err)
	}

	i := 0

	walter, err := entity.NewChef("Walter White", "Master Chef", "Descubra o mundo dos sabores através dos olhos e das mãos de um chef apaixonado pela culinária. Nossa busca incessante pela excelência nos leva a explorar a arte da gastronomia de maneira única, criando experiências culinárias que cativam o paladar e encantam os sentidos (Jesse, we need to cook!)", "")
	if err != nil {
		panic(err)
	}

	jessy, err := entity.NewChef("Jesse Pinkman", "Sous Chef", "Você está olhando para o coração pulsante da nossa cozinha. Nosso Sous Chef desempenha um papel crucial na criação das delícias culinárias que saem de nossa cozinha e é o guardião da qualidade e da excelência. (Are we in the chef business, or the money business?)", "")
	if err != nil {
		panic(err)
	}
	gus, err := entity.NewChef("Gustavo Fring", "Cook", "Conheça o mais recente talento na nossa equipe culinária. Este chef emergente traz consigo uma paixão inabalável pela gastronomia e um desejo de impressionar os comensais com suas habilidades únicas na cozinha. (I cook In plain sight, same as you.)", "")
	if err != nil {
		panic(err)
	}
	chefs := []schemas.Chef{schemas.Chef{gorm.Model{}, *walter}, schemas.Chef{gorm.Model{}, *jessy}, schemas.Chef{gorm.Model{}, *gus}}

	db := config.GetDb()
	for i < 3 {

		chefs[i].ProfilePic = fmt.Sprintf("uploads/%s", files[i].Name())
		db.Save(&chefs[i])
		logger.Debug(chefs[i])
		i += 1
	}
}

func Seed() {
	SeedChef()
	SeedFood()
}

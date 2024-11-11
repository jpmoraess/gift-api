package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
	"github.com/jpmoraess/gift-api/internal/infra"
)

func main() {
	app := fiber.New(fiber.Config{})

	giftRepo := infra.NewGiftRepositoryAdapter()

	createGift := usecase.NewCreateGift(giftRepo)

	app.Post("/gifts", func(c *fiber.Ctx) error {
		input := new(usecase.CreateGiftInput)
		if err := c.BodyParser(input); err != nil {
			return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
		}

		output, err := createGift.Execute(input)
		if err != nil {
			log.Fatal(err.Error())
		}
		return c.JSON(output)
	})

	app.Listen(":8080")
}

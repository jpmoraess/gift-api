package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/jpmoraess/gift-api/db/sqlc"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
	"github.com/jpmoraess/gift-api/internal/infra"
	"log"
)

func main() {
	dbConfig, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(pool)

	// repo
	giftRepo := infra.NewGiftRepositoryAdapter(store)

	// usecase
	createGift := usecase.NewCreateGift(giftRepo)

	app := fiber.New(fiber.Config{})

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

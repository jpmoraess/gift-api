package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/gift-api/config"
	db "github.com/jpmoraess/gift-api/db/sqlc"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
	"github.com/jpmoraess/gift-api/internal/infra/factory"
	"github.com/jpmoraess/gift-api/internal/infra/gateway"
	"github.com/jpmoraess/gift-api/internal/infra/persistence"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
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

	// store
	store := db.NewStore(pool)

	// payment gateway
	asaasGateway := gateway.NewAsaas(&config, &http.Client{})

	// factory
	paymentProcessorFactory := factory.NewPaymentProcessorFactory(asaasGateway)

	// processor
	paymentProcessor := paymentProcessorFactory.CreatePaymentProcessor()

	// repository
	giftRepository := persistence.NewGiftRepositoryAdapter(store)
	transactionRepository := persistence.NewTransactionRepositoryAdapter(store)

	// usecase
	createGift := usecase.NewCreateGift(giftRepository)
	processPayment := usecase.NewProcessPayment(paymentProcessor, transactionRepository)

	app := fiber.New(fiber.Config{})

	handleCreateGift(app, createGift)
	handleProcessPayment(app, processPayment)

	app.Listen(":8080")
}

func handleCreateGift(app *fiber.App, createGift *usecase.CreateGift) {
	app.Post("/gifts", func(c *fiber.Ctx) (err error) {
		input := new(usecase.CreateGiftInput)
		if err = c.BodyParser(input); err != nil {
			return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
		}

		output, err := createGift.Execute(context.Background(), input)
		if err != nil {
			log.Fatal(err.Error())
		}
		return c.JSON(output)
	})
}

func handleProcessPayment(app *fiber.App, processPayment *usecase.ProcessPayment) {
	app.Post("/transactions", func(c *fiber.Ctx) (err error) {
		input := new(usecase.ProcessPaymentInput)
		if err = c.BodyParser(input); err != nil {
			return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
		}

		output, err := processPayment.Execute(context.Background(), input)
		if err != nil {
			log.Fatal(err.Error())
		}
		return c.JSON(output)
	})
}

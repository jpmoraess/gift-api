package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/gift-api/config"
	db "github.com/jpmoraess/gift-api/db/sqlc"
	_ "github.com/jpmoraess/gift-api/docs"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
	"github.com/jpmoraess/gift-api/internal/infra/factory"
	"github.com/jpmoraess/gift-api/internal/infra/gateway"
	"github.com/jpmoraess/gift-api/internal/infra/handlers"
	"github.com/jpmoraess/gift-api/internal/infra/persistence"
	"github.com/jpmoraess/gift-api/token"
	"log"
	"net/http"
)

// @title			I-GIFT
// @version		1.0
// @description	I-GIFT is a platform for you to give gifts to your friends and family
// @termsOfService	http://swagger.io/terms/
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

	// token maker
	tokenMaker, err := token.NewPasetoMaker([]byte(config.SymmetricKey))
	if err != nil {
		log.Fatal(err)
	}

	// store
	store := db.NewStore(pool)

	// payment gateway
	asaasPaymentGateway := gateway.NewAsaasPaymentGateway(&config, &http.Client{})

	// factory
	paymentProcessorFactory := factory.NewPaymentProcessorFactory(asaasPaymentGateway)

	// processor
	paymentProcessor := paymentProcessorFactory.CreatePaymentProcessor()

	// repository
	giftRepository := persistence.NewGiftRepositoryAdapter(store)
	userRepository := persistence.NewUserRepositoryAdapter(store)
	transactionRepository := persistence.NewTransactionRepositoryAdapter(store)

	// usecase
	createGift := usecase.NewCreateGift(giftRepository)
	createUser := usecase.NewCreateUser(userRepository)
	generateToken := usecase.NewGenerateToken(tokenMaker, userRepository)
	processPayment := usecase.NewProcessPayment(paymentProcessor, transactionRepository)

	// fiber
	app := fiber.New(fiber.Config{})

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// handlers
	giftHandler := handlers.NewGiftHandler(createGift)
	userHandler := handlers.NewUserHandler(createUser)
	tokenHandler := handlers.NewTokenHandler(generateToken)
	transactionHandler := handlers.NewTransactionHandler(processPayment)

	// routes
	app.Post("/auth/token", func(c *fiber.Ctx) error {
		return tokenHandler.GenerateToken(c)
	})

	app.Post("/v1/gifts", func(c *fiber.Ctx) error {
		return giftHandler.CreateGift(c)
	})

	app.Post("/v1/users", func(c *fiber.Ctx) error {
		return userHandler.CreateUser(c)
	})

	app.Post("/v1/transactions", func(c *fiber.Ctx) error {
		return transactionHandler.ProcessPayment(c)
	})

	app.Listen(":8080")
}

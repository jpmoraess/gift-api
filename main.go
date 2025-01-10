package main

import (
	"context"
	"log"
	"net/http"

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
	"github.com/jpmoraess/gift-api/internal/infra/storage"
	"github.com/jpmoraess/gift-api/token"
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
	asaasGateway := gateway.NewAsaasGateway(&config, &http.Client{})

	// factory
	chargeGeneratorFactory := factory.NewChargeGeneratorFactory(asaasGateway)

	// chain
	chargeGenerator := chargeGeneratorFactory.CreateChargeGeneratorChain()

	// repository
	fileRepository := storage.NewFileRepository(store)
	userRepository := persistence.NewUserRepositoryAdapter(store)
	transactionRepository := persistence.NewTransactionRepositoryAdapter(store)

	// usecase
	createUser := usecase.NewCreateUser(userRepository)
	generateToken := usecase.NewGenerateToken(tokenMaker, userRepository)
	generateCharge := usecase.NewGenerateCharge(chargeGenerator, transactionRepository)

	// storage
	localStorage := storage.NewLocalStorage(config.FilePath)

	// services
	fileService := storage.NewFileService(localStorage, fileRepository)

	// fiber
	app := fiber.New(fiber.Config{})

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// handlers
	userHandler := handlers.NewUserHandler(createUser)
	fileHandler := handlers.NewFileHandler(fileService)
	tokenHandler := handlers.NewTokenHandler(generateToken)
	transactionHandler := handlers.NewTransactionHandler(generateCharge)

	// routes
	RegisterRoutes(app, userHandler, fileHandler, tokenHandler, transactionHandler)

	app.Listen(":8080")
}

func RegisterRoutes(
	app *fiber.App,
	userHandler *handlers.UserHandler,
	fileHandler *handlers.FileHandler,
	tokenHandler *handlers.TokenHandler,
	transactionHandler *handlers.TransactionHandler,
) {
	app.Post("/auth/token", func(c *fiber.Ctx) error {
		return tokenHandler.GenerateToken(c)
	})

	app.Post("/v1/users", func(c *fiber.Ctx) error {
		return userHandler.CreateUser(c)
	})

	app.Post("/v1/transactions", func(c *fiber.Ctx) error {
		return transactionHandler.CreateTransaction(c)
	})

	app.Post("/v1/files", func(c *fiber.Ctx) error {
		return fileHandler.Upload(c)
	})

	app.Get("/v1/files/:id", func(c *fiber.Ctx) error {
		return fileHandler.Download(c)
	})

	app.Delete("/v1/files/:id", func(c *fiber.Ctx) error {
		return fileHandler.Delete(c)
	})
}

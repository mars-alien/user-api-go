package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/recover"
    _ "github.com/lib/pq"

    "github.com/mars-alien/user-api-go/config"
    "github.com/mars-alien/user-api-go/internal/handler"
    "github.com/mars-alien/user-api-go/internal/logger"
    "github.com/mars-alien/user-api-go/internal/repository"
    "github.com/mars-alien/user-api-go/internal/routes"
    "github.com/mars-alien/user-api-go/internal/service"
)

func main() {
    if err := logger.InitLogger(); err != nil {
        log.Fatal("Failed to initialize logger:", err)
    }
    defer logger.Sync()

    cfg, err := config.LoadConfig()
    if err != nil {
        logger.Log.Fatal("Failed to load config")
    }

    db, err := sql.Open("postgres", cfg.GetDBConnectionString())
    if err != nil {
        logger.Log.Fatal("Failed to connect to database")
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        logger.Log.Fatal("Failed to ping database")
    }

    logger.Log.Info("Connected to database successfully")

    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    app := fiber.New(fiber.Config{
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }
            return c.Status(code).JSON(fiber.Map{
                "error": err.Error(),
            })
        },
    })

    app.Use(recover.New())
    app.Use(cors.New())

    routes.SetupRoutes(app, userHandler)

    addr := fmt.Sprintf(":%s", cfg.ServerPort)
    logger.Log.Info("Server starting on " + addr)
    
    if err := app.Listen(addr); err != nil {
        logger.Log.Fatal("Failed to start server")
    }
}
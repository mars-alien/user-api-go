package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/mars-alien/user-api-go/internal/handler"
    "github.com/mars-alien/user-api-go/internal/middleware"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
    app.Use(middleware.RequestLogger())
    
    api := app.Group("/users")
    
    api.Post("/", userHandler.CreateUser)
    api.Get("/:id", userHandler.GetUser)
    api.Get("/", userHandler.ListUsers)
    api.Put("/:id", userHandler.UpdateUser)
    api.Delete("/:id", userHandler.DeleteUser)
}
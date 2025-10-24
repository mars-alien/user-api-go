package handler

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/go-playground/validator/v10"
    "go.uber.org/zap"

    "github.com/mars-alien/user-api-go/internal/logger"
    "github.com/mars-alien/user-api-go/internal/models"
    "github.com/mars-alien/user-api-go/internal/service"
)

type UserHandler struct {
    service   service.UserService
    validator *validator.Validate
}

func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{
        service:   service,
        validator: validator.New(),
    }
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var req models.CreateUserRequest
    
    if err := c.BodyParser(&req); err != nil {
        logger.Log.Error("Failed to parse request body", zap.Error(err))
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := h.validator.Struct(req); err != nil {
        logger.Log.Error("Validation failed", zap.Error(err))
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    user, err := h.service.CreateUser(c.Context(), req)
    if err != nil {
        logger.Log.Error("Failed to create user", zap.Error(err))
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create user",
        })
    }

    logger.Log.Info("User created successfully", zap.Int32("user_id", user.ID))
    return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    user, err := h.service.GetUser(c.Context(), int32(id))
    if err != nil {
        logger.Log.Error("Failed to get user", zap.Error(err), zap.Int("user_id", id))
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    return c.JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
    page := c.QueryInt("page", 1)
    pageSize := c.QueryInt("page_size", 10)

    users, err := h.service.ListUsers(c.Context(), page, pageSize)
    if err != nil {
        logger.Log.Error("Failed to list users", zap.Error(err))
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to list users",
        })
    }

    return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    var req models.UpdateUserRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := h.validator.Struct(req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    user, err := h.service.UpdateUser(c.Context(), int32(id), req)
    if err != nil {
        logger.Log.Error("Failed to update user", zap.Error(err), zap.Int("user_id", id))
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update user",
        })
    }

    logger.Log.Info("User updated successfully", zap.Int("user_id", id))
    return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    if err := h.service.DeleteUser(c.Context(), int32(id)); err != nil {
        logger.Log.Error("Failed to delete user", zap.Error(err), zap.Int("user_id", id))
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete user",
        })
    }

    logger.Log.Info("User deleted successfully", zap.Int("user_id", id))
    return c.SendStatus(fiber.StatusNoContent)
}
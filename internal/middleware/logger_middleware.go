package middleware

import (
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "go.uber.org/zap"

    "github.com/mars-alien/user-api-go/internal/logger"
)

func RequestLogger() fiber.Handler {
    return func(c *fiber.Ctx) error {
        requestID := uuid.New().String()
        c.Set("X-Request-ID", requestID)
        
        start := time.Now()
        
        err := c.Next()
        
        duration := time.Since(start)
        
        logger.Log.Info("Request completed",
            zap.String("request_id", requestID),
            zap.String("method", c.Method()),
            zap.String("path", c.Path()),
            zap.Int("status", c.Response().StatusCode()),
            zap.Duration("duration", duration),
        )
        
        return err
    }
}
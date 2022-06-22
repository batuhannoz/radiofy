package middleware

import "github.com/gofiber/fiber/v2"

type LogService interface {
	//
}

type Log struct {
	LogService
}

func NewLogMiddleware(logService LogService) *Log {
	return &Log{
		logService,
	}
}
func (log *Log) Log(c *fiber.Ctx) error {
	c.Next()
	return nil
}

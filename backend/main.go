package main

import (
	"backend/internal"
	"backend/internal/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	conf, err := config.New(".config", appEnv)
	if err != nil {
		fmt.Println(err)
	}
	conf.Print()

	app := fiber.New()
	app.Use(cors.New())

	internal.RegisterRoutes(conf, app)

	err = app.Listen(conf.Server.Port)
	if err != nil {
		fmt.Println(err)
	}
}

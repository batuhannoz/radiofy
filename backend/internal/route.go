package internal

import (
	"backend/internal/config"
	appHandler "backend/internal/handler/app"
	"backend/internal/middleware"
	"backend/internal/service"
	"backend/internal/store"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(config *config.Config, app *fiber.App) {
	db := store.NewMySQL(config.MySQL)

	logStore := store.NewLogStore(db)
	userStore := store.NewUserStore(db)
	clubStore := store.NewClubStore(db)
	chatStore := store.NewChatStore(db)

	logService := service.NewLogService(logStore)
	userService := service.NewUserService(userStore)
	clubService := service.NewClubService(clubStore)
	chatService := service.NewChatService(chatStore)

	appHandler := appHandler.NewHandler(userService, clubService, chatService)

	logMiddleware := middleware.NewLogMiddleware(logService)
	jwtMiddleware := middleware.NewJWTMiddleware(userService, clubService, chatService)

	routes := app.Group("/", logMiddleware.Log)
	{
		routes.Get("/login", appHandler.Login)

		routesWithJWT := routes.Group("/", jwtMiddleware.Authorize)
		{
			routesWithJWT.Get("/clubs", appHandler.Clubs)
			routesWithJWT.Delete("/logout", appHandler.Logout)
		}
	}
}

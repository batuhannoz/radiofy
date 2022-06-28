package internal

import (
	"backend/internal/config"
	appHandler "backend/internal/handler/app"
	"backend/internal/middleware"
	"backend/internal/service"
	"backend/internal/store"
	"github.com/gofiber/fiber/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

func RegisterRoutes(config *config.Config, app *fiber.App) {
	db := store.NewMySQL(config.MySQL)

	var (
		auth = spotifyauth.New(
			spotifyauth.WithClientID(config.Spotify.ClientID),
			spotifyauth.WithClientSecret(config.Spotify.ClientSecret),
			spotifyauth.WithRedirectURL(config.Spotify.RedirectURI),
			spotifyauth.WithScopes(
				spotifyauth.ScopeStreaming,
				spotifyauth.ScopeUserReadEmail,
				spotifyauth.ScopePlaylistModifyPrivate,
				spotifyauth.ScopeUserTopRead,
				spotifyauth.ScopeUserReadRecentlyPlayed,
				spotifyauth.ScopeUserReadPlaybackState,
				spotifyauth.ScopeUserReadCurrentlyPlaying,
				spotifyauth.ScopeUserReadPrivate,
				spotifyauth.ScopeUserModifyPlaybackState,
				spotifyauth.ScopePlaylistReadCollaborative),
		)
	)

	logStore := store.NewLogStore(db)
	userStore := store.NewUserStore(db)
	clubStore := store.NewClubStore(db)
	chatStore := store.NewChatStore(db)

	logService := service.NewLogService(logStore)
	userService := service.NewUserService(userStore, config)
	clubService := service.NewClubService(clubStore)
	chatService := service.NewChatService(chatStore)
	authService := service.NewAuthService(auth, config, userService)

	appHandler := appHandler.NewHandler(userService, clubService, chatService, authService)

	logMiddleware := middleware.NewLogMiddleware(logService)
	jwtMiddleware := middleware.NewJWTMiddleware(userService, clubService, chatService)

	routes := app.Group("/", logMiddleware.Log)
	{
		routes.Get("/auth_url", appHandler.AuthURL)
		routes.Get("/complete_auth", appHandler.CompleteAuth)

		routesWithJWT := routes.Group("/", jwtMiddleware.Authorize)
		{
			routesWithJWT.Get("/clubs", appHandler.Clubs)
			routesWithJWT.Post("/create_club", appHandler.CreateClub)
		}
	}
}

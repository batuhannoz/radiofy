package app

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserService interface {
	//
}

type ClubService interface {
	Clubs() *ClubsResponse
	CreateClub() *CreateClubResponse
}

type ChatService interface {
	//
}

type AuthService interface {
	AuthURL() *AuthURLResponse
	CompleteAuth(ctx *fiber.Ctx) *TokenResponse
}

type AppHandler struct {
	userService UserService
	clubService ClubService
	ChatService ChatService
	authService AuthService
}

func NewHandler(userService UserService, clubService ClubService, chatService ChatService, authService AuthService) *AppHandler {
	return &AppHandler{
		userService,
		clubService,
		chatService,
		authService,
	}
}

func (app *AppHandler) AuthURL(c *fiber.Ctx) error {
	url := app.authService.AuthURL()
	return c.Status(http.StatusOK).JSON(url)
}

func (app *AppHandler) CompleteAuth(c *fiber.Ctx) error {
	token := app.authService.CompleteAuth(c)
	return c.Status(http.StatusOK).JSON(token)
}

func (app *AppHandler) Clubs(c *fiber.Ctx) error {
	//Clubs := app.clubService.Clubs()
	//return c.Status(http.StatusOK).JSON(Clubs)
	return c.Status(http.StatusOK).JSON("pong")
}

func (app *AppHandler) CreateClub(c *fiber.Ctx) error {
	Club := app.clubService.CreateClub()
	return c.Status(http.StatusOK).JSON(Club)
}

package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"net/http"
)

type UserService interface {
	//
}

type ClubService interface {
	Clubs() *[]ClubsResponse
	CreateClub(ctx *fiber.Ctx, ClubRequest CreateClubRequest) *CreateClubResponse
	ListenSong(c *websocket.Conn) error
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
	Clubs := app.clubService.Clubs()
	return c.Status(http.StatusOK).JSON(Clubs)
}

func (app *AppHandler) CreateClub(c *fiber.Ctx) error {
	var ClubRequest CreateClubRequest
	err := c.BodyParser(&ClubRequest)
	if err != nil {
		fmt.Println(err)
		return err
	}
	Club := app.clubService.CreateClub(c, ClubRequest)
	return c.Status(http.StatusOK).JSON(Club)
}

func (app *AppHandler) ListenSong(c *websocket.Conn) {
	err := app.clubService.ListenSong(c)
	if err != nil {
		fmt.Println(err)
	}
}

func (app *AppHandler) ChangeSong(ctx *fiber.Ctx) error {
	// change song
	ctx.Status(http.StatusOK)
	return nil
}

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
	CurrentSong(ctx *fiber.Ctx) *PlaybackResponse
	ChangeSong(ctx *fiber.Ctx, SongRequest *PlaybackRequest) error
	Listener(ws *websocket.Conn)
}

type ChatService interface {
	GlobalChat(ws *websocket.Conn)
	ClubChat(ws *websocket.Conn)
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

func (app *AppHandler) CurrentSong(c *fiber.Ctx) error {
	songPlayback := app.clubService.CurrentSong(c)
	return c.Status(http.StatusOK).JSON(songPlayback)
}

func (app *AppHandler) ChangeSong(ctx *fiber.Ctx) error {
	var SongRequest PlaybackRequest
	err := ctx.BodyParser(&SongRequest)
	if err != nil {
		fmt.Println(err)
	}
	err = app.clubService.ChangeSong(ctx, &SongRequest)
	if err != nil {
		fmt.Println(err)
	}
	ctx.Status(http.StatusOK)
	return nil
}

func (app *AppHandler) Listener(ws *websocket.Conn) {
	app.clubService.Listener(ws)
}

func (app *AppHandler) GlobalChat(ws *websocket.Conn) {
	app.ChatService.GlobalChat(ws)
}

func (app *AppHandler) ClubChat(ws *websocket.Conn) {
	app.ChatService.ClubChat(ws)
}

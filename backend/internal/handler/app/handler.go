package app

import (
	"backend/internal/store/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type UserService interface {
	Login(request model.User) *model.User
}

type ClubService interface {
	//
}

type ChatService interface {
	//
}

type AppHandler struct {
	userService UserService
	clubService ClubService
	ChatService ChatService
}

//
type user struct {
	SpotifyID string `json:"spotify_id"`
	Mail      string `json:"mail"`
	Country   string `json:"country"`
	Product   string `json:"product"`
}

//

func NewHandler(userService UserService, clubService ClubService, chatService ChatService) *AppHandler {
	return &AppHandler{
		userService,
		clubService,
		chatService,
	}
}

func (app *AppHandler) Login(c *fiber.Ctx) error {
	usr := user{}
	err := c.BodyParser(&usr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr)
	date := time.Now()
	user := model.User{
		Id:         0,
		CreateDate: date,
		SpotifyID:  usr.SpotifyID,
		Mail:       usr.Mail,
		Country:    usr.Country,
		Product:    usr.Product,
	}
	a := app.userService.Login(user)
	return c.Status(http.StatusOK).JSON(a)
}

func (app *AppHandler) Clubs(c *fiber.Ctx) error {
	return nil
}

func (app *AppHandler) Logout(c *fiber.Ctx) error {
	return nil
}

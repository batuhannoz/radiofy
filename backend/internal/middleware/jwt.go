package middleware

import "github.com/gofiber/fiber/v2"

type UserService interface {
	//
}

type ClubService interface {
	//
}

type ChatService interface {
	//
}

type JWT struct {
	UserService
	ClubService
	ChatService
}

func NewJWTMiddleware(userService UserService, clubService ClubService, chatService ChatService) *JWT {
	return &JWT{
		userService,
		clubService,
		chatService,
	}
}

func (jwt *JWT) Authorize(c *fiber.Ctx) error {
	return nil
}

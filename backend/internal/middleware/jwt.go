package middleware

import (
	"backend/internal/config"
	"backend/internal/store/model"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
)

type UserService interface {
	GetUserActiveUserLogon(userId uint64) (*model.User, *model.UserLogon)
}

type ClubService interface {
	//
}

type ChatService interface {
	//
}

type JWT struct {
	Config *config.Config
	UserService
	ClubService
	ChatService
}

func NewJWTMiddleware(config *config.Config, userService UserService, clubService ClubService, chatService ChatService) *JWT {
	return &JWT{
		config,
		userService,
		clubService,
		chatService,
	}
}

func (j *JWT) Authorize(c *fiber.Ctx) error {
	token := string(c.Request().Header.Peek("Authorization"))
	if token == "" {
		token = c.Query("token")
	}
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON("no token provided")
	}
	userId, errToken := j.getUserId(token)
	if errToken != nil {
		return c.Status(http.StatusUnauthorized).JSON("token is not valid")
	}

	user, userLogon := j.UserService.GetUserActiveUserLogon(userId)
	if userLogon.Token != token || userLogon.IsLogout == true {
		return c.Status(http.StatusUnauthorized).JSON("token valid but not active")
	}

	c.Locals("user", user)
	c.Next()
	return nil
}

func (j *JWT) validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.Config.JWT.Secret), nil
	})
}

func (j *JWT) getUserId(token string) (uint64, error) {
	jwtToken, errToken := j.validateToken(token)
	if errToken != nil {
		_ = fmt.Errorf("Token Error: %v\n", errToken.Error())
		return 0, errors.New(errToken.Error())
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userId, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["userID"]), 10, 64)
	return userId, nil
}

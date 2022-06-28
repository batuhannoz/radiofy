package service

import (
	"backend/internal/config"
	"backend/internal/functions"
	"backend/internal/handler/app"
	"backend/internal/store/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	spotifyAuth "github.com/zmb3/spotify/v2/auth"
	"net/http"
	"net/url"
	"time"
)

type AuthService struct {
	Config      *config.Config
	Spotify     *spotifyAuth.Authenticator
	UserService *UserService
}

func NewAuthService(spotify *spotifyAuth.Authenticator, config *config.Config, userService *UserService) *AuthService {
	return &AuthService{
		config,
		spotify,
		userService,
	}
}

func (auth *AuthService) AuthURL() *app.AuthURLResponse {
	url := auth.Spotify.AuthURL(functions.EncodeToString(6))
	return &app.AuthURLResponse{
		url,
	}
}

func (auth *AuthService) CompleteAuth(ctx *fiber.Ctx) *app.TokenResponse {
	var r http.Request
	r.URL = &url.URL{
		RawQuery: "code=" + ctx.Query("code") + "&" + "state=" + ctx.Query("state") + "&" + "error=" + ctx.Query("error"),
	}
	spotifyToken, err := auth.Spotify.Token(ctx.Context(), ctx.Query("state"), &r)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	user := auth.UserService.SaveUser(spotifyToken)
	radiofyToken := auth.UserService.generateJWTToken(user.Id)
	auth.UserService.addUserLogon(model.UserLogon{
		Id:         0,
		UserID:     user.Id,
		CreateDate: time.Now(),
		Token:      radiofyToken,
		IsLogout:   false,
	})
	return &app.TokenResponse{
		AccessToken:  spotifyToken.AccessToken,
		RefreshToken: spotifyToken.RefreshToken,
		ExpiryTime:   spotifyToken.Expiry,
		RadiofyToken: radiofyToken,
	}
}

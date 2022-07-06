package service

import (
	"backend/internal/config"
	"backend/internal/store"
	"backend/internal/store/model"
	"fmt"
	spotifyUser "github.com/batuhannoz/spotigo/user"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"time"
)

type UserStore interface {
	SaveUser(user model.User) *model.User
	AddUserLogon(userLogon *model.UserLogon) *model.UserLogon
	GetUserActiveUserLogon(userId uint64) *model.UserLogon
	GetUserById(userId uint64) *model.User
}

type UserService struct {
	Config    *config.Config
	UserStore UserStore
}

func NewUserService(userStore *store.UserStore, config *config.Config) *UserService {
	return &UserService{
		config,
		userStore,
	}
}

func (u *UserService) SaveUser(token *oauth2.Token) *model.User {
	user, err := spotifyUser.GetCurrentUsersProfile(spotifyUser.GetCurrentUsersProfileInput{Token: "Bearer " + token.AccessToken})
	if err != nil {
		fmt.Println(err)
	}
	imgLength := len(user.Images)
	if imgLength == 0 {
		return u.UserStore.SaveUser(model.User{
			Id:          0,
			DisplayName: user.DisplayName,
			CreateDate:  time.Now(),
			Image:       "",
			SpotifyID:   user.Id,
			Mail:        user.Email,
			Country:     user.Country,
			Product:     user.Product,
		})
	}
	return u.UserStore.SaveUser(model.User{
		Id:          0,
		DisplayName: user.DisplayName,
		CreateDate:  time.Now(),
		Image:       user.Images[0].Url,
		SpotifyID:   user.Id,
		Mail:        user.Email,
		Country:     user.Country,
		Product:     user.Product,
	})
}

func (u *UserService) generateJWTToken(userID uint64) string {
	claims := jwt.MapClaims{
		"userID":  userID,
		"expires": time.Now().Add(time.Hour * u.Config.JWT.Expire).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(u.Config.JWT.Secret))
	if err != nil {
		// TODO: throw exception
	}
	return t
}

func (u *UserService) addUserLogon(userLogon model.UserLogon) {
	u.UserStore.AddUserLogon(&userLogon)
}

func (u *UserService) GetUserActiveUserLogon(userId uint64) (*model.User, *model.UserLogon) {
	return u.UserStore.GetUserById(userId), u.UserStore.GetUserActiveUserLogon(userId)
}

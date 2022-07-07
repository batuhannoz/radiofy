package service

import (
	"backend/internal/config"
	"backend/internal/handler/app"
	"backend/internal/store"
	"backend/internal/store/model"
	"fmt"
	spotifyUser "github.com/batuhannoz/spotigo/user"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"time"
)

type ActiveListener struct {
	position    int
	albumId     string
	displayName string
	image       string
	artistName  string
	songName    string
	songID      string
	ws          *websocket.Conn
}

type UserStore interface {
	SaveUser(user model.User) *model.User
	AddUserLogon(userLogon *model.UserLogon) *model.UserLogon
	GetUserActiveUserLogon(userId uint64) *model.UserLogon
	GetUserById(userId uint64) *model.User
}

type UserService struct {
	Config          *config.Config
	UserStore       UserStore
	ActiveListeners map[uint64]*ActiveListener
}

func NewUserService(userStore *store.UserStore, config *config.Config) *UserService {
	var ListenerClients = make(map[uint64]*ActiveListener)
	return &UserService{
		config,
		userStore,
		ListenerClients,
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

func (u *UserService) CurrentListeners(ws *websocket.Conn) {
	user := ws.Locals("user").(*model.User)
	u.ActiveListeners[user.Id] = &ActiveListener{
		albumId:     "",
		position:    0,
		displayName: user.DisplayName,
		image:       user.Image,
		artistName:  "",
		songName:    "",
		songID:      "",
		ws:          ws,
	}
	u.updateListeners()

	u.ActiveListeners[user.Id].ws.SetCloseHandler(func(code int, text string) error {
		delete(u.ActiveListeners, user.Id)
		u.updateListeners()
		return nil
	})
	for {
		var request app.CurrentListenersRequest
		err := ws.ReadJSON(&request)
		if err != nil {
			fmt.Println(err)
			return
		}
		u.ActiveListeners[user.Id].songID = request.SongID
		u.ActiveListeners[user.Id].songName = request.SongName
		u.ActiveListeners[user.Id].artistName = request.ArtistName
		u.ActiveListeners[user.Id].position = request.Position
		u.ActiveListeners[user.Id].albumId = request.AlbumID
		u.updateListeners()
	}
}

func (u *UserService) updateListeners() {
	var response []app.CurrentListenersResponse
	for _, listener := range u.ActiveListeners {
		response = append(response, app.CurrentListenersResponse{
			AlbumID:     listener.albumId,
			Position:    listener.position,
			DisplayName: listener.displayName,
			Image:       listener.image,
			ArtistName:  listener.artistName,
			SongName:    listener.songName,
			SongID:      listener.songID,
		})
	}
	for _, listener := range u.ActiveListeners {
		err := listener.ws.WriteJSON(response)
		if err != nil {
			fmt.Println(err)
		}
	}
}

package service

import (
	"backend/internal/handler/app"
	"backend/internal/store/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strconv"
	"time"
)

type ClubStore interface {
	Clubs() *[]model.Club
	CreateClub(club *model.Club) *model.Club
	ClubByOwnerID(ownerID uint64) *model.Club
	ClubListeners(ClubID int) *[]model.Listener
	AddListener(listener *model.Listener) *model.Listener
	CurrentSongByClubID(clubID int) *model.ClubSong
	ChangeClubSong(clubSong *model.ClubSong) *model.ClubSong
	DeactivateListener(listenerID uint64)
}

type Client struct {
	id       string
	ip       string
	username string
	ws       *websocket.Conn
}

var Clients = make(map[uint64]*Client)

type ClubService struct {
	ClubStore
}

func NewClubService(clubStore ClubStore) *ClubService {
	return &ClubService{
		clubStore,
	}
}

func (c *ClubService) Clubs() *[]app.ClubsResponse {
	clubs := c.ClubStore.Clubs()
	var clubsResponse []app.ClubsResponse
	for _, b := range *clubs {
		clubsResponse = append(clubsResponse, app.ClubsResponse{
			ID:          b.Id,
			Image:       b.Image,
			Name:        b.Name,
			Description: b.Description,
		})
	}
	return &clubsResponse
}

func (c *ClubService) CreateClub(ctx *fiber.Ctx, ClubRequest app.CreateClubRequest) *app.CreateClubResponse {
	// TODO check if user is trying to open more than one club
	user := ctx.Locals("user").(*model.User)
	CreatedClub := c.ClubStore.CreateClub(&model.Club{
		Id:          0,
		OwnerID:     user.Id,
		Name:        ClubRequest.Name,
		IsActive:    true,
		Description: ClubRequest.Description,
		Image:       ClubRequest.Image,
	})
	return &app.CreateClubResponse{
		Code:        strconv.Itoa(CreatedClub.Id),
		Name:        CreatedClub.Name,
		Image:       CreatedClub.Image,
		Description: CreatedClub.Description,
	}
}

func (c *ClubService) CurrentSong(ctx *fiber.Ctx) *app.PlaybackResponse {
	clubID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		fmt.Println(err)
	}
	club := c.ClubStore.CurrentSongByClubID(clubID)
	return &app.PlaybackResponse{
		AlbumID:    club.AlbumID,
		Position:   club.Position,
		SongID:     club.SongID,
		SongName:   club.SongName,
		ArtistName: club.ArtistName,
		Image:      club.Image,
		ProgressMS: club.ProgressMS,
	}
}

func (c *ClubService) ChangeSong(ctx *fiber.Ctx, request *app.PlaybackRequest) error {
	user := ctx.Locals("user").(*model.User)
	club := c.ClubStore.ClubByOwnerID(user.Id)
	if strconv.Itoa(club.Id) != ctx.Params("id") {
		fmt.Println("wrong request")
	}
	c.ClubStore.ChangeClubSong(&model.ClubSong{
		Id:         0,
		ClubID:     club.Id,
		AlbumID:    request.AlbumID,
		Position:   request.Position,
		SongID:     request.SongID,
		SongName:   request.SongName,
		ArtistName: request.ArtistName,
		Image:      request.Image,
		ProgressMS: request.ProgressMS,
	})
	c.UpdateListenersPlayback(user.Id, request)
	return nil
}

func (c *ClubService) Listener(ws *websocket.Conn) {
	user := ws.Locals("user").(*model.User)
	c.ClubStore.AddListener(&model.Listener{
		Id:       0,
		UserID:   user.Id,
		ClubID:   ws.Params("id"),
		IsActive: true,
	})
	newClient := Client{
		id:       "",
		ip:       ws.RemoteAddr().String(),
		username: "",
		ws:       ws,
	}
	Clients[user.Id] = &newClient
	Clients[user.Id].ws.SetCloseHandler(func(code int, text string) error {
		c.ClubStore.DeactivateListener(user.Id)
		delete(Clients, user.Id)
		return nil
	})
	newClient.StartListening(user.Id)
}

func (client *Client) StartListening(userID uint64) {
	for {
		var msg string
		err := Clients[userID].ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(5000)
	}
}

func (c *ClubService) UpdateListenersPlayback(ownerID uint64, request *app.PlaybackRequest) {
	club := c.ClubStore.ClubByOwnerID(ownerID)
	listeners := c.ClubStore.ClubListeners(club.Id)
	for _, listener := range *listeners {
		err := Clients[listener.UserID].ws.WriteJSON(&request)
		if err != nil {
			fmt.Println(err)
		}
	}
}

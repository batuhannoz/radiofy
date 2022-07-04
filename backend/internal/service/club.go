package service

import (
	"backend/internal/handler/app"
	"backend/internal/store/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strconv"
)

type ClubStore interface {
	Clubs() *[]model.Club
	CreateClub(club *model.Club) *model.Club
	ClubByOwnerID(ownerID int) *model.Club
	ClubListeners(ClubID int) *[]model.Listener
	AddListener(listener *model.Listener) *model.Listener
	CurrentSongByClubID(clubID int) *model.ClubSong
	ChangeClubSong(clubSong *model.ClubSong) *model.ClubSong
}

type ClubService struct {
	ClubStore
	writerByUserID map[uint64]websocket.Conn
}

func NewClubService(clubStore ClubStore) *ClubService {
	writer := make(map[uint64]websocket.Conn)
	return &ClubService{
		clubStore,
		writer,
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
	club := c.ClubStore.ClubByOwnerID(int(user.Id))
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
	return nil
}

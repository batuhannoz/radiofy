package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
)

var ClubWithCodes = make(map[string]Club)

type Playback struct {
	SongID     string `json:"songID"`
	SongName   string `json:"songName"`
	ArtistName string `json:"artistName"`
	Image      string `json:"image"`
	ProgressMS int    `json:"progressMS"`
}

type Club struct {
	Code        string `json:"code"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var WsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Listener(ctx *gin.Context) {
	// TODO Connect to users with websocket and send the song the leader is listening to
}

func ClubLeader(ctx *gin.Context) {
	var playback Playback

	err := json.NewDecoder(ctx.Request.Body).Decode(&playback)
	if err != nil {
		fmt.Println(err)
	}
	// TODO send incoming playback to other listeners
}

func CreateClub(ctx *gin.Context) {
	var c Club

	err := json.NewDecoder(ctx.Request.Body).Decode(&c)
	if err != nil {
		fmt.Println(err)
	}
	code := EncodeToString(6)
	c.Code, ClubWithCodes[code] = code, c
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	ctx.String(http.StatusOK, string(data))
}

func ActiveClubsList(ctx *gin.Context) {
	data, err := json.Marshal(ClubWithCodes)
	if err != nil {
		fmt.Println(err)
	}
	ctx.String(http.StatusOK, string(data))
}

func main() {
	club1 := Club{
		Code:        "111111",
		Image:       "https://i.scdn.co/image/993cdbae6bf9c21aa653ada2d153c8f05f1b842e",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	club2 := Club{
		Code:        "222222",
		Image:       "https://i.scdn.co/image/ab67616d00001e026ca5c90113b30c3c43ffb8f4",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	club3 := Club{
		Code:        "333333",
		Image:       "https://i.scdn.co/image/ab67706c0000bebb443715c10bcde6c10f733874",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	club4 := Club{
		Code:        "444444",
		Image:       "https://i.scdn.co/image/ab67616d00001e02d9b35d1c4d15c9de88b754a7",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	club5 := Club{
		Code:        "555555",
		Image:       "https://i.scdn.co/image/ab67616d00001e0259ac0a9acb7bf0d315301152",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	ClubWithCodes["111111"] = club1
	ClubWithCodes["222222"] = club2
	ClubWithCodes["333333"] = club3
	ClubWithCodes["444444"] = club4
	ClubWithCodes["555555"] = club5

	r := gin.Default()
	r.Use(cors.Default())
	//
	// Path that gives instantly active clubs
	r.GET("/club/list", ActiveClubsList)
	//
	// The path the club leader uses to change songs
	r.GET("/club/leader/:id", ClubLeader)
	//
	// Websocket path to which listeners connect
	r.GET("/club/:id", Listener)
	//
	// The path the club leader uses to create a club and get their own party/club code
	r.POST("/create/club", CreateClub)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}

// EncodeToString Generate random string
var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// Mysql DB Connection With GORM
/*
func (UserInfo) TableName() string {
	return "user_info"
}

type UserInfo struct {
	Id             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	DisplayName    string `gorm:"type:varchar(50)" json:"display_name"`
	Mail           string `gorm:"type:varchar(50)" json:"mail"`
	Country        string `gorm:"type:varchar(50)" json:"country"`
	ProfilePicture string `gorm:"type:varchar(100)" json:"profile_picture"`

func main() {
	dbUser := ""
	dbPass := ""
	dbHost := ""
	dbPort := ""
	dbName := ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	r.GET("/play", StartResume)
	r.GET("/pause", Pause)
	r.GET("/volume", Volume)
	r.GET("/adduser", AddUser)
	r.Run(":3000")
}

 func AddUser(gin *gin.Context) {
	a := gin.Request.URL.Query()
	token := a["token"][0]
	realtoken := "Bearer " + token
	c, err := spotify.GetCurrentUsersProfile(spotify.GetCurrentUsersProfileInput{Token: realtoken})
	if err != nil {
		fmt.Println(err)
	}

	user_info := UserInfo{
		DisplayName:    c.DisplayName,
		Mail:           c.Email,
		Country:        c.Country,
		ProfilePicture: c.Images[0].Url,
	}
	db.Save(&user_info)
}
*/

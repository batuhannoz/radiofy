package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

/*type Message struct {
	ClubName        string `json:"ClubName"`
	ClubDescription string `json:"ClubDescription"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		// check the http.Request
		// make sure it's OK to access
		return true
	}
	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return
	}

	defer wsConn.Close()

	// event loop
	for {
		var msg Message

		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}
		fmt.Printf("Club Name: %s\n", msg.ClubName)
		fmt.Printf("Club Description: %s\n", msg.ClubDescription)
		var ReceivedMSG Message
		ReceivedMSG.ClubName = msg.ClubName
		ReceivedMSG.ClubDescription = msg.ClubDescription
		err = wsConn.WriteJSON(ReceivedMSG)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SendMessage(msg string) {
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}*/

func createClub(ctx *gin.Context) {

}

type Club struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func activeClubsList(ctx *gin.Context) {
	club1 := &Club{
		Image:       "https://i.scdn.co/image/993cdbae6bf9c21aa653ada2d153c8f05f1b842e",
		Name:        "Science Club",
		Description: "listening to science podcasts",
	}
	club2 := &Club{
		Image:       "https://i.scdn.co/image/ab67616d00001e026ca5c90113b30c3c43ffb8f4",
		Name:        "Eminem Club",
		Description: "listening to Eminem songs",
	}
	club3 := &Club{
		Image:       "https://i.scdn.co/image/ab67706c0000bebb443715c10bcde6c10f733874",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	club4 := &Club{
		Image:       "https://i.scdn.co/image/ab67616d00001e02d9b35d1c4d15c9de88b754a7",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	club5 := &Club{
		Image:       "https://i.scdn.co/image/ab67616d00001e0259ac0a9acb7bf0d315301152",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	}
	data, err := json.Marshal([]*Club{club1, club2, club3, club4, club5})
	if err != nil {
		fmt.Println(err)
	}
	ctx.String(200, string(data))
}

func main() {
	r := gin.Default()
	r.GET("/club/list", CORSMiddleware(), activeClubsList)
	r.GET("/create/club", CORSMiddleware(), createClub)
	r.Run(":3000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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

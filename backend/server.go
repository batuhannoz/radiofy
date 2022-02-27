package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Message struct {
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
}

func main() {
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		WsEndpoint(c.Writer, c.Request)
	})
	r.Run(":3000")
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

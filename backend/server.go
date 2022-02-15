package main

import (
	"fmt"
	spotify "github.com/batuhannoz/spotigo/user"
	"github.com/gin-gonic/gin"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (UserInfo) TableName() string {
	return "user_info"
}

type UserInfo struct {
	Id             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	DisplayName    string `gorm:"type:varchar(50)" json:"display_name"`
	Mail           string `gorm:"type:varchar(50)" json:"mail"`
	Country        string `gorm:"type:varchar(50)" json:"country"`
	ProfilePicture string `gorm:"type:varchar(100)" json:"profile_picture"`
}

const (
	SetVolumeUrl   = "https://api.spotify.com/v1/me/player/volume"
	StartResumeUrl = "https://api.spotify.com/v1/me/player/play"
	PauseUrl       = "https://api.spotify.com/v1/me/player/pause"
	AuthName       = "Authorization"
)

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

func Pause(*gin.Context) {
	url := PauseUrl
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPut, url, nil)
	//req.Header.Set(AuthName, Token)
	resp, _ := client.Do(req)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func StartResume(*gin.Context) {
	url := StartResumeUrl
	client := http.Client{}

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//req.Header.Set("Authorization", Token)
	response, _ := client.Do(req)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func Volume(gin *gin.Context) {
	a := gin.Request.URL.Query()
	b := a["volume"][0]
	url := SetVolumeUrl
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//req.Header.Set(AuthName, Token)
	q := req.URL.Query()
	q.Add("volume_percent", b)
	req.URL.RawQuery = q.Encode()
	response, _ := client.Do(req)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
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

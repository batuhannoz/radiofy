package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	SetVolumeUrl   = "https://api.spotify.com/v1/me/player/volume"
	StartResumeUrl = "https://api.spotify.com/v1/me/player/play"
	PauseUrl       = "https://api.spotify.com/v1/me/player/pause"
	AuthName       = "Authorization"
	Token          = "Bearer <token>"
)

func main() {
	r := gin.Default()
	r.GET("/play", StartResume)
	r.GET("/pause", Pause)
	r.GET("/volume", Volume)
	r.Run()
}

func Pause(*gin.Context) {
	url := PauseUrl
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPut, url, nil)
	req.Header.Set(AuthName, Token)
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
	req.Header.Set("Authorization", Token)
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
	req.Header.Set(AuthName, Token)
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

package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Chat struct {
	Text string `json:"text"`
	Time string `json:"time"`
}

func ChatHandler(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("%s", bytes)

	var chat Chat

	json.Unmarshal(bytes, &chat)
	log.Printf("text: %s, time: %s", chat.Text, chat.Time)

	c.Status(http.StatusOK)
}

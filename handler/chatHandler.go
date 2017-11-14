package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Chat struct {
	Text       string `json:"text"`
	Time       string `json:"time"`
	SpeakerId  int64  `json:"speaker-id"`
	ChatRoomId int64  `json:"chat-room-id"`
}

func ChatHandler(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	//log.Printf("%s", bytes)

	var chat Chat

	json.Unmarshal(bytes, &chat)
	log.Printf("text: %s, time: %s, speakerId: %d, chatRoomId: %d", chat.Text, chat.Time, chat.SpeakerId, chat.ChatRoomId)

	c.Status(http.StatusOK)
}

package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Chat struct {
	Text       string `json:"text"`
	Time       string `json:"time"`
	SpeakerId  string `json:"speaker-id"`
	ChatRoomId string `json:"chat-room-id"`
}

func ChatHandler(c *gin.Context) {
	text, _ := c.GetPostForm("text")
	time, _ := c.GetPostForm("time")
	speakerId, _ := c.GetPostForm("speaker-id")
	chatRoomId, _ := c.GetPostForm("chat-room-id")
	var chat = Chat{text, time, speakerId, chatRoomId}

	log.Printf("text: %s, time: %s, speakerId: %s, chatRoomId: %s\n", chat.Text, chat.Time, chat.SpeakerId, chat.ChatRoomId)

	c.Status(http.StatusOK)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func ChatHandler(c *gin.Context) {
	buff, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("%s", buff)

	c.Status(http.StatusOK)
}

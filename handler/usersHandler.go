package handler

import (
	"log"

	"net/http"

	"bytes"

	"github.com/gin-gonic/gin"
)

func UsersHandler(c *gin.Context) {
	log.Printf("%#v", c)

	body := c.Request.Body

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	s := buf.String() + "\n"

	c.String(http.StatusOK, s)
}

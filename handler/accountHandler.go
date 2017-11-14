package handler

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

type Account struct {
	Uuid string
}

func AccountHandler(c *gin.Context) {

	switch c.Request.Method {
	case "GET":
		getHandler(c)
		break
	case "POST":
		postHandler(c)
		break
	}
}

func getHandler(c *gin.Context) {
	uuid, _ := c.GetQuery("uuid")
	account := Account{uuid}
	log.Printf("uuid: %s", account.Uuid)

	stubJson := `
{
    "data": {
        "user-id": 10,
		"display-name": "kazu"
    }
}
`
	c.String(http.StatusOK, stubJson)
}

func postHandler(c *gin.Context) {
	println("post")
}

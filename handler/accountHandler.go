package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountHandler(c *gin.Context) {

	switch c.Request.Method {
	case "GET":
		getHandler(c)
	case"POST":
		postHandler(c)
	}
}

func getHandler(c *gin.Context) {
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

}
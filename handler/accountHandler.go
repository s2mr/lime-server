package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountHandler(c *gin.Context) {

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

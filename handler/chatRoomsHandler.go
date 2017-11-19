package handler

import (
	"log"

	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func ChatRoomsHandler(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("%s\n", bytes)

	stubJson := `
{
    "data": {
        "chatRooms": [{
                "id": 1,
                "user": {
                    "id": 1,
                    "name": "たろー",
                    "screenName": "たろー",
                    "statusText": "ステータス1"
                },
                "currentText": "この本まじ面白くて、読み出すと本当止まらないんだよね笑今度かそっか？😉"
            },
            {
                "id": 2,
                "user": {
                    "id": 2,
                    "name": "対話BOT",
                    "screenName": "対話BOT",
                    "statusText": "対話BOT"
                },
                "currentText": ""
            }
        ]
    }
}
`
	c.String(http.StatusOK, stubJson)
}

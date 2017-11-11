package handler

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func FriendsHandler(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("%s", bytes)

	stubJson := `
{
    "data": {
        "users": [
            {
                "id": 1,
                "name": "名前1",
                "screenName": "スクリーンネーム1",
                "statusText": "ステータス1",
            },
            {
                "id": 2,
                "name": "名前2",
                "screenName": "スクリーンネーム2",
                "statusText": "ステータス2",
            },
            {
                "id": 3,
                "name": "名前3",
                "screenName": "スクリーンネーム3",
                "statusText": "ステータス3",
            },
        ]
    }
}
`


	c.String(http.StatusOK, stubJson)
}

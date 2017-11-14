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
                    "name": "名前1",
                    "screenName": "スクリーンネーム1",
                    "statusText": "ステータス1"
                },
                "currentText": "テキスト1"
            },
            {
                "id": 2,
                "user": {
                    "id": 2,
                    "name": "名前2",
                    "screenName": "スクリーンネーム2",
                    "statusText": "ステータス2"
                },
                "currentText": "テキスト2"
            },
            {
                "id": 3,
                "user": {
                    "id": 3,
                    "name": "名前3",
                    "screenName": "スクリーンネーム3",
                    "statusText": "ステータス3"
                },
                "currentText": "テキスト3"
            },
            {
                "id": 4,
                "user": {
                    "id": 4,
                    "name": "名前4",
                    "screenName": "スクリーンネーム4",
                    "statusText": "ステータス4"
                },
                "currentText": "テキスト4"
            }
        ]
    }
}
`
	c.String(http.StatusOK, stubJson)
}

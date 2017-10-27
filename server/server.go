package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shimokp/lime-server/handler"
)

func Init() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	v1 := router.Group("/v1")
	{
		v1.GET("/", handler.IndexHandler)
	}

	router.Run(":8080")

}

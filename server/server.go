package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shimokp/lime-server/handler"
)

func Init() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	root := router.Group("/")
	{
		root.GET("/", handler.IndexHandler)
	}

	v1 := router.Group("/v1")
	{
		v1.POST("/users", handler.UsersHandler)
	}

	port := os.Getenv("PORT")

	router.Run(":" + port)

}

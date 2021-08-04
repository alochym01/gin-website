package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router Gin Engine
func Router() *gin.Engine {
	router := gin.Default()

	// Using anonymous function
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Using normal function
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", PostAlbums)

	return router
}

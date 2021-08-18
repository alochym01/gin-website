package router

import (
	"net/http"

	"github.com/alochym01/gin-website/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	albumH := &handler.AlbumHandler{}

	// Using anonymous function
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Using normal function
	// Albums Handler
	router.GET("/albums", albumH.Get)
	router.GET("/albums/:id", albumH.GetByID)
	router.POST("/albums", albumH.Post)
	router.PUT("/albums/:id", albumH.UpdateByID)

	return router

}

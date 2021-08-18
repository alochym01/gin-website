package router

import (
	"net/http"

	"github.com/alochym01/gin-website/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	// Using anonymous function
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	albumH := handler.NewAlbumHandler()
	// Albums Handler
	// Using normal function
	router.GET("/albums", albumH.Index)
	router.GET("/albums/:id", albumH.Show)
	router.POST("/albums", albumH.Create)
	router.PUT("/albums/:id", albumH.Update)
	router.DELETE("/albums/:id", albumH.Delete)

	return router

}

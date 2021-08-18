package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alochym01/gin-website/handler"
	"github.com/gin-gonic/gin"
)

// Router is returned *gin.Engine
func Router() *gin.Engine {
	// Disable log's color
	gin.DisableConsoleColor()

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\" \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
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

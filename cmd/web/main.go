package main

import (
	"github.com/alochym01/gin-website/api"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 117.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	router := api.Router()
	// router := gin.Default()

	// // Using anonymous function
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// // Using normal function
	// router.GET("/albums", api.getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	// A expvar handler for gin framework
	router.Run()
}

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, entity.Albums)
// }

// // postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
// 	var newAlbum entity.Album

// 	// Call BindJSON to bind the received JSON to newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid Params",
// 		})
// 		return
// 	}

// 	// Add the new album to the slice.
// 	entity.Albums = append(entity.Albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range entity.Albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

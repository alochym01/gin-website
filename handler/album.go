package handler

import (
	"database/sql"
	"net/http"

	"github.com/alochym01/gin-website/models"
	"github.com/alochym01/gin-website/repository"
	"github.com/gin-gonic/gin"
)

// AlbumHandler is Album Handler
type AlbumHandler struct {
	repoAlbum repository.AlbumRepo
}

// NewAlbumHandler is return AlbumHandler
func NewAlbumHandler() *AlbumHandler {
	return &AlbumHandler{
		repoAlbum: models.Album{},
	}
}

// Index list of all albums as JSON in DB.
func (e AlbumHandler) Index(c *gin.Context) {
	var album models.Album

	albums, err := album.Get()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, "Server Error")
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// Show a record of albums as JSON in DB.
func (e AlbumHandler) Show(c *gin.Context) {
	var album models.Album
	ID, _ := c.Params.Get("id")

	// Single-Row Queries
	// a, _ := models.Album.GetByID(ID) ==> fail not sure why
	record, err := album.GetByID(ID)

	if err != nil {
		// No record in database
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, "Not found")
			return
		} else {
			// check server err
			c.IndentedJSON(http.StatusBadGateway, "Server Error")
			return
		}
	}

	c.IndentedJSON(http.StatusOK, record)
}

// Create an album from JSON received in the request body.
func (e AlbumHandler) Create(c *gin.Context) {
	var newAlbum models.RequestAlbum
	var album models.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Params",
		})
		return
	}

	err := album.Create(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "server error")
		return
	}
	c.IndentedJSON(http.StatusCreated, "OK")
}

// Update a record of albums as JSON into DB.
// Request body should contains:
// title
// artist
// price
func (e AlbumHandler) Update(c *gin.Context) {
	var updateRequest models.RequestAlbum
	var album models.Album

	// Try to get ID from request
	ID, err1 := c.Params.Get("id")
	if err1 == false {
		c.IndentedJSON(http.StatusBadRequest, "ID require")
		return
	}

	// Check record exist in DB
	_, err := album.GetByID(ID)

	if err != nil {
		// No record in database
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusBadRequest, "Try again")
			return
		} else {
			// check server err
			c.IndentedJSON(http.StatusBadGateway, "Server Error")
			return
		}
	}

	// convert request body in json format to updateAlbum object
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update record exist in DB
	err = album.Update(updateRequest.Title, updateRequest.Artist, updateRequest.Price, ID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Server error")
		return
	}
	c.IndentedJSON(http.StatusAccepted, "OK")
}

// Delete a record of albums in DB.
func (e AlbumHandler) Delete(c *gin.Context) {
	var album models.Album
	ID, _ := c.Params.Get("id")

	err := album.Delete(ID)

	if err != nil {
		// No record in database
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, "Not found")
			return
		} else {
			// check server err
			c.IndentedJSON(http.StatusBadGateway, "Server Error")
			return
		}
	}

	c.IndentedJSON(http.StatusOK, "OK")
}

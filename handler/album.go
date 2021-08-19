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
	albumRepo repository.AlbumRepo
}

// NewAlbumHandler is return AlbumHandler
func NewAlbumHandler() *AlbumHandler {
	return &AlbumHandler{
		albumRepo: models.Album{},
	}
}

// Index list of all albums as JSON in DB.
func (al AlbumHandler) Index(c *gin.Context) {
	// Using album repository interface
	albums, err := al.albumRepo.Get()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, "Server Error")
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// Show a record of albums as JSON in DB.
func (al AlbumHandler) Show(c *gin.Context) {
	id := c.Param("id")

	// Single-Row Queries
	// Using album repository interface
	record, err := al.albumRepo.GetByID(id)

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
func (al AlbumHandler) Create(c *gin.Context) {
	var newAlbum models.RequestAlbum

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Params",
		})
		return
	}

	// Using album repository interface
	err := al.albumRepo.Create(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Server error")
		return
	}

	c.IndentedJSON(http.StatusCreated, "OK")
}

// Update a record of albums as JSON into DB.
// Request body should contains:
// title
// artist
// price
func (al AlbumHandler) Update(c *gin.Context) {
	var updateRequest models.RequestAlbum

	// Try to get ID from request
	ID, err1 := c.Params.Get("id")
	if err1 == false {
		c.IndentedJSON(http.StatusBadRequest, "ID require")
		return
	}

	// Check record exist in DB
	// Using album repository interface
	_, err := al.albumRepo.GetByID(ID)

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
	// Using album repository interface
	err = al.albumRepo.Update(updateRequest.Title, updateRequest.Artist, updateRequest.Price, ID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Server error")
		return
	}

	c.IndentedJSON(http.StatusAccepted, "OK")
}

// Delete a record of albums in DB.
func (al AlbumHandler) Delete(c *gin.Context) {
	ID, _ := c.Params.Get("id")

	// Using album repository interface
	err := al.albumRepo.Delete(ID)

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

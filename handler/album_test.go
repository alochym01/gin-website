package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alochym01/gin-website/config"
	"github.com/alochym01/gin-website/models"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

// https://github.com/gin-gonic/gin/issues/2833 - cannot use gin.CreateTestContext()

func setupRouter() *gin.Engine {
	router := gin.Default()
	albumH := NewAlbumHandler()

	router.GET("/albums", albumH.Index)
	router.GET("/albums/:id", albumH.Show)
	router.POST("/albums", albumH.Create)
	router.PUT("/albums/:id", albumH.Update)
	router.DELETE("/albums/:id", albumH.Delete)

	return router
}

func TestAlbumHandler(t *testing.T) {

	config.DB = config.SqliteConn("../foo.db")
	defer config.DB.Close()

	models.PreparesqliteDB(config.DB)

	e := setupRouter()

	t.Run("Test AlbumHandler Index Method", func(t *testing.T) {
		route := "/albums"

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", route, nil)

		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Test AlbumHandler Show Method", func(t *testing.T) {
		route := "/albums/1"

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", route, nil)

		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Test AlbumHandler Create Method", func(t *testing.T) {
		body := `{ "Title":  "title", "Artist": "artist", "Price":  7.99 }`
		route := "/albums"

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", route, strings.NewReader(string(body)))

		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Test AlbumHandler Update Method", func(t *testing.T) {
		body := `{ "Title":  "title", "Artist": "artist", "Price":  7.99 }`
		route := "/albums/1"

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("PUT", route, strings.NewReader(string(body)))

		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusAccepted, w.Code)
	})

	t.Run("Test AlbumHandler Delete Method", func(t *testing.T) {
		route := "/albums/1"

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", route, nil)

		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

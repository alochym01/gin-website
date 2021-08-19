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

func TestAlbumHandler(t *testing.T) {
	config.DB = config.SqliteConn("../foo.db")
	defer config.DB.Close()

	models.PreparesqliteDB(config.DB)
	albumH := NewAlbumHandler()

	t.Run("Test AlbumHandler Index Method", func(t *testing.T) {
		route := "/albums"

		w := httptest.NewRecorder()

		c, r := gin.CreateTestContext(w)

		r.GET(route, albumH.Index)

		c.Request, _ = http.NewRequest("GET", route, nil)

		r.ServeHTTP(w, c.Request)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Test AlbumHandler Show Method", func(t *testing.T) {
		w := httptest.NewRecorder()

		c, e := gin.CreateTestContext(w)

		route := "/albums/1"

		e.GET(route, albumH.Show)

		c.Request, _ = http.NewRequest("GET", route, nil)

		e.ServeHTTP(w, c.Request)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Test AlbumHandler Create Method", func(t *testing.T) {
		body := `{ "Title":  "title", "Artist": "artist", "Price":  7.99 }`
		w := httptest.NewRecorder()

		c, e := gin.CreateTestContext(w)

		route := "/albums"

		e.POST(route, albumH.Create)

		c.Request, _ = http.NewRequest("POST", route, strings.NewReader(string(body)))

		e.ServeHTTP(w, c.Request)

		assert.Equal(t, 201, w.Code)
	})

	t.Run("Test AlbumHandler Update Method", func(t *testing.T) {
		body := `{ "Title":  "title", "Artist": "artist", "Price":  7.99 }`
		w := httptest.NewRecorder()

		c, e := gin.CreateTestContext(w)

		route := "/albums/1"

		e.PUT(route, albumH.Update)

		c.Request, _ = http.NewRequest("PUT", route, strings.NewReader(string(body)))

		e.ServeHTTP(w, c.Request)

		assert.Equal(t, 201, w.Code)
	})

	t.Run("Test AlbumHandler Delete Method", func(t *testing.T) {
		w := httptest.NewRecorder()

		c, e := gin.CreateTestContext(w)

		route := "/albums/1"

		e.DELETE(route, albumH.Delete)

		c.Request, _ = http.NewRequest("DELETE", route, nil)

		e.ServeHTTP(w, c.Request)

		assert.Equal(t, 201, w.Code)
	})

}

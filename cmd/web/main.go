package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	os.Remove("foo.db")

	// sqlite section
	db, _ := sql.Open("sqlite3", "./foo.db")
	defer db.Close()

	// Create Album table
	Album := `CREATE TABLE albums (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"artist" TEXT,
		"price" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(Album) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Album table created")

	// Prepare to insert value
	stmt, _ := db.Prepare("INSERT INTO albums(title, artist, price) values(?,?,?)")

	for _, v := range albums {
		fmt.Println(v.ID)
		fmt.Println(v.Artist)
		fmt.Println(v.Title)
		fmt.Println(v.Price)

		// fmt.Println(stmt)
		_, _ = stmt.Exec(v.Title, v.Artist, v.Price)
	}

	router := gin.Default()

	// Using anonymous function
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Using normal function
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run()
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Params",
		})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

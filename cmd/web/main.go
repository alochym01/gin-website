package main

import (
	"fmt"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Env struct {
	db *sql.DB
}

type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Remove sqlite db file
	os.Remove("foo.db")

	// sqlite section
	db, _ := sql.Open("sqlite3", "./foo.db")
	defer db.Close()

	env := &Env{db: db}

	prepareDB(db)

	router := gin.Default()

	// Using anonymous function
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Using normal function
	router.GET("/albums", env.getAlbums)
	router.GET("/albums/:id", env.getAlbumsByID)
	router.POST("/albums", env.postAlbums)
	router.PUT("/albums/:id", env.updateAlbumsByID)

	router.Run()
}

// getAlbums responds with the list of all albums as JSON.
func (e *Env) getAlbums(c *gin.Context) {
	var (
		// Create empty album
		record = album{}
		// Create empty slice album
		albums = []album{}
	)

	// Mutiple-Rows Queries
	rows, _ := e.db.Query("SELECT id, title, artist, price FROM albums")

	for rows.Next() {
		err := rows.Scan(&record.ID, &record.Title, &record.Artist, &record.Price)
		if err != nil {
			// No record in database
			if err == sql.ErrNoRows {
				fmt.Println(err)
				c.IndentedJSON(http.StatusNotFound, "Not found")
				return
			} else {
				// check server err
				c.IndentedJSON(http.StatusBadGateway, "Server Error")
				return

			}
		}

		// append record to albums slice
		albums = append(albums, record)
	}

	// close rows query
	defer rows.Close()

	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbums responds with the a record of albums as JSON.
func (e *Env) getAlbumsByID(c *gin.Context) {
	var record = album{}

	ID, _ := c.Params.Get("id")

	// sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("SELECT id, title, artist, price FROM albums where id=%s", ID)

	// Single-Row Queries
	err := e.db.QueryRow(sqlstmt).Scan(&record.ID, &record.Title, &record.Artist, &record.Price)

	if err != nil {
		// No record in database
		if err == sql.ErrNoRows {
			fmt.Println(err)
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

// postAlbums adds an album from JSON received in the request body.
func (e *Env) postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Params",
		})
		return
	}
	// sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("INSERT INTO albums(id, title, artist, price) VALUES(%d, \"%s\", \"%s\", %f)",
		newAlbum.ID,
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)
	fmt.Println(sqlstmt)

	// Execute SQL Statements
	_, err := e.db.Exec(sqlstmt)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "server error")
		return
	}
	c.IndentedJSON(http.StatusCreated, "OK")
}

// getAlbums responds with the a record of albums as JSON.
func (e *Env) updateAlbumsByID(c *gin.Context) {
	type updateAlbum struct {
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	}
	var updateRequest updateAlbum
	var record = album{}

	ID, _ := c.Params.Get("id")

	// sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("SELECT id, title, artist, price FROM albums where id=%s", ID)

	// Single-Row Queries
	err := e.db.QueryRow(sqlstmt).Scan(&record.ID, &record.Title, &record.Artist, &record.Price)

	if err != nil {
		// No record in database
		if err == sql.ErrNoRows {
			fmt.Println(err)
			c.IndentedJSON(http.StatusNotFound, "Not found")
			return
		} else {
			// check server err
			c.IndentedJSON(http.StatusBadGateway, "Server Error")
			return
		}
	}

	// convert request body in json format to updateAlbum object, request body contains:
	// title
	// artist
	// price
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// sqlstmt - Avoid SQL Injection Attack
	updatestmt := fmt.Sprintf("UPDATE albums SET title=\"%s\", artist=\"%s\", price=%f where id=%s",
		updateRequest.Title,
		updateRequest.Artist,
		updateRequest.Price,
		ID,
	)
	fmt.Println(updatestmt)

	// Execute SQL Statements
	_, err = e.db.Exec(updatestmt)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "server error")
		return
	}
	c.IndentedJSON(http.StatusAccepted, "OK")
}

// prepareDB create mock data.
func prepareDB(db *sql.DB) {
	// Create Album table
	// SQL Statement for Create Table
	sqltable := fmt.Sprint("CREATE TABLE albums ('id' integer NOT NULL PRIMARY KEY AUTOINCREMENT,'title' TEXT,'artist' TEXT, 'price' TEXT);")

	// Execute SQL Statements
	_, err := db.Exec(sqltable)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Album table created")

	for _, v := range albums {
		sqlstmt := fmt.Sprintf("INSERT INTO albums(title, artist, price) VALUES(\"%s\", \"%s\", %f)",
			v.Title,
			v.Artist,
			v.Price,
		)
		// fmt.Println(v.ID, v.Artist, v.Title, v.Price)
		// Execute SQL Statements
		db.Exec(sqlstmt) // Execute SQL Statements
	}
}

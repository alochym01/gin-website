package models

import (
	"fmt"

	"github.com/alochym01/gin-website/config"
)

// Album models
type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// RequestAlbum models
type RequestAlbum struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Get Method get a record from DB
func (al Album) Get() ([]Album, error) {
	var (
		// Create empty album
		record = Album{}
		// Create empty slice album
		albums = []Album{}
	)

	// sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("SELECT id, title, artist, price FROM albums")

	fmt.Println(sqlstmt)

	rows, _ := config.DB.Query(sqlstmt)

	for rows.Next() {
		err := rows.Scan(&record.ID, &record.Title, &record.Artist, &record.Price)
		if err != nil {
			// check server err
			return albums, err
		}

		// append record to albums slice
		albums = append(albums, record)
	}

	// close rows query
	defer rows.Close()

	return albums, nil
}

// GetByID Method get a record from DB
func (al Album) GetByID(id string) (Album, error) {
	var record = Album{}

	// sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("SELECT id, title, artist, price FROM albums where id=%s", id)

	fmt.Println(sqlstmt)

	// Single-Row Queries
	err := config.DB.QueryRow(sqlstmt).Scan(&record.ID, &record.Title, &record.Artist, &record.Price)

	return record, err
}

// Create Method a record into DB
func (al Album) Create(title string, artist string, price float64) error {

	// sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("INSERT INTO albums(title, artist, price) VALUES(\"%s\", \"%s\", %f)",
		title,
		artist,
		price,
	)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	_, err := config.DB.Exec(sqlstmt)
	if err != nil {
		return err
	}
	return nil
}

// Update Method a record into DB
func (al Album) Update(title string, artist string, price float64, id string) error {
	// sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("UPDATE albums SET title=\"%s\", artist=\"%s\", price=%f where id=%s",
		title,
		artist,
		price,
		id,
	)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	_, err := config.DB.Exec(sqlstmt)
	if err != nil {
		return err
	}
	return nil
}

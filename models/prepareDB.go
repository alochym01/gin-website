package models

import (
	"database/sql"
	"fmt"
)

// PrepareMysqlDB create mock data.
func PrepareMysqlDB(db *sql.DB) {
	// Create Album table
	// SQL Statement for Create Table
	sqltable := fmt.Sprint("CREATE TABLE albums (id INT NOT NULL AUTO_INCREMENT,title TEXT,artist TEXT, price TEXT, PRIMARY KEY (id));")
	sqldroptable := fmt.Sprint("DROP TABLE albums;")
	db.Exec(sqldroptable)

	// Execute SQL Statements
	_, err := db.Exec(sqltable)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Album table created")

	// albums slice to seed record album data.
	var albums = []Album{
		{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

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

// PreparesqliteDB create mock data.
func PreparesqliteDB(db *sql.DB) {
	// Create Album table
	// SQL Statement for Create Table
	sqltable := fmt.Sprint("CREATE TABLE albums (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,title TEXT,artist TEXT, price TEXT);")

	// Execute SQL Statements
	_, err := db.Exec(sqltable)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Album table created")

	// albums slice to seed record album data.
	var albums = []Album{
		{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

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

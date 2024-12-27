package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to the PostgreSQL database
	connStr := "user=postgres dbname=learninggo password=sa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()


	rows2, err := db.Query(" INSERT INTO list VALUES (8, 'BumrahPaji') ")
	if err != nil {
		panic(err)
	}

	defer rows2.Close()

	// Perform a sample query
	rows, err := db.Query("SELECT * FROM list")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Iterate through the results
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	connStr := "postgres://postgres:password@localhost:5432/db_fooddelivery?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var results []sql.NullString

	userRow, err := db.Query("SELECT first_name FROM customer")
	if err == sql.ErrNoRows {
		// Check for empty result set
		fmt.Println("NO ROWS !!!")
	} else if err != nil {
		panic(err)
	}

	for userRow.Next() {
		username := sql.NullString{}
		err = userRow.Scan(&username)
		if err != nil {
			panic(err)
		}

		results = append(results, username)
	}

	for _, name := range results {
		v, _ := name.Value()
		fmt.Println("name.Value() - ", v)
		fmt.Println("name.Valid   - ", name.Valid)
		fmt.Println("---")
	}

}

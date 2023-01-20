package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
/*
In this code, the user input is not properly validated or 
sanitized before being used in the SQL query. 
An attacker could provide a malicious input, such as test' OR 1=1 -- in 
order to modify the query and potentially gain unauthorized access to the database.
*/

func main() {
	db, _ := sql.Open("sqlite3", "file:test.db")
	defer db.Close()

	userInput := "test' OR 1=1 --"
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", userInput)

	rows, _ := db.Query(query)
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		rows.Scan(&id, &username)
		fmt.Println(id, username)
	}
}

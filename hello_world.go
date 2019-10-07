package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var username string
var password string

const (
	host   = "localhost"
	port   = 5432
	dbname = "db1"
)

func init() {
	fmt.Scanln(&username, &password)
}

func executeQuery(db *sql.DB) {
	query := `SELECT * from table_for_one;`
	rows, _ := db.Query(query)
	for rows.Next() {
		var (
			userId   int
			username string
			password string
			email    string
		)

		rows.Scan(&userId, &username, &password, &email)

		fmt.Println(userId, username, password, email)
	}

}

func main() {
	fmt.Println("Hello " + username + "!" + " Password: " + password)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	executeQuery(db)
}

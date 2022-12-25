package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to  database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, age FROM users where id=$1")
	if err != nil {
		log.Fatal("con't prepare query all users statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id, age int
	var name string

	err = row.Scan(&id, &name, &age)

	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	fmt.Println("one raw", id, name, age)

}
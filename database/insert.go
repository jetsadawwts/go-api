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

	raw :=  db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", "Wongwit", 21)
	var id int
	err = raw.Scan(&id)
	if err != nil {
		log.Fatal("can't insert data", err)
	} 
	fmt.Println("insert todo success id : ", id)	

}
package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1:3306)/many_data")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	_, err = dbConnection.Exec(`CREATE TABLE IF NOT EXISTS bugs (id INT, bug_count INT, date_reported DATE)`)
	if err != nil {
		log.Fatal(err)
	}

	insertDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	for i := 1; i < 1000000; i++ {
		insertDate = insertDate.Add(time.Duration(24) * time.Hour)
		_, err := dbConnection.Exec(`INSERT INTO bugs (id, bug_count, date_reported) VALUES (?, ?, ?)`, i, rand.Intn(100), insertDate.Format("2006-01-02"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1:3307)/many_data")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	_, err = dbConnection.Exec(`CREATE TABLE IF NOT EXISTS bugs (id INT, bug_id INT, date_reported DATE)`)
	if err != nil {
		log.Fatal(err)
	}

	insertDate := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	for i := 1; i < 100000; i++ {
		insertDate = insertDate.Add(time.Duration(24) * time.Hour)
		_, err := dbConnection.Exec(`INSERT INTO bugs (id, bug_id, date_reported) VALUES (?, ?, ?)`, i, i, insertDate.Format("2006-01-02"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
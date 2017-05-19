package db

import (
	"database/sql"
	"fmt"
	//		"time"

	_ "github.com/mattn/go-sqlite3"
)

var dbfd *sql.DB

func init() {
	fmt.Println("init database")
	db, err := sql.Open("sqlite3", "../data.db")
	checkErr(err)
	dbfd = db
	fmt.Println("init databas done")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

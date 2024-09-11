package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const ConnectionString = "postgres://postgres:secret@localhost:5432/task_management#?sslmode=disable"

func DbInstance() *sql.DB {
	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Fialed To Ping Database", err)
	}
	fmt.Println("Connected Successfully")

	return db

}

var Client *sql.DB = DbInstance()

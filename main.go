package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	id        int64
	title     string
	content   string
	published bool
	author    User
	authorId  int64
}

type User struct {
	id    int64
	name  string
	email string
	posts []Post
}

func Insert(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS test_user (
		id SERIAL NOT NULL PRIMARY KEY,
		name VARCHAR(10))`)
		if err != nil {
			return err
		}

		_, err = db.Exec("INSERT INTO test_user (name) VALUES ('tom')")
		if err != nil {
			return err
		}
		return nil
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5433 user=admin password=admin dbname=db sslmode=disable")
	if err != nil {
		log.Fatal("OpenError:", err)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("PingError:", err)
	}
}
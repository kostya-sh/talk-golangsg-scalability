package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/lib/pq"
)

const table = "temperature"

func createTable(db *sql.DB) {
	if _, err := db.Exec("drop table if exists " + table); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("create table " + table + " (time varchar(30), city varchar(30), value varchar(30))"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "postgres://u:p@localhost/d?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := txn.Prepare(pq.CopyIn(table, "time", "city", "value"))
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		tokens := strings.Split(s.Text(), ",")
		if len(tokens) != 3 {
			log.Fatal("invalid line: ", s.Text())
		}

		_, err = stmt.Exec(tokens[0], tokens[1], tokens[2])
		if err != nil {
			log.Fatal("exec: ", err)
		}
	}

	if _, err = stmt.Exec(); err != nil {
		log.Fatal("final exec: ", err)
	}
	// if err = stmt.Close(); err != nil {
	// 	log.Fatal("close: ", err)
	// }
	if err = txn.Commit(); err != nil {
		log.Fatal("commit: ", err)
	}
}

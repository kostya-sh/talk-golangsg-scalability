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

func handleLine(st *sql.Stmt, line string) {
	tokens := strings.Split(line, ",")
	if len(tokens) != 3 {
		log.Fatal("invalid line: ", line)
	}
	if _, err := st.Exec(tokens[0], tokens[1], tokens[2]); err != nil { // HL
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

	// START_MAIN OMIT
	txn, err := db.Begin() // HL
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := txn.Prepare(pq.CopyIn(table, "time", "city", "value")) // HL
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		handleLine(stmt, s.Text()) // HL
	}
	if s.Err() != nil {
		log.Fatal("scan: ", s.Err())
	}
	if _, err = stmt.Exec(); err != nil { // HL
		log.Fatal("final exec: ", err)
	}
	if err = txn.Commit(); err != nil { // HL
		log.Fatal("commit: ", err)
	}
	// END_MAIN OMIT
}

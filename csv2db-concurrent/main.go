package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"
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

// START_HL OMIT
func handleLine(st *sql.Stmt, line string) {
	tokens := strings.Split(line, ",")
	if len(tokens) != 3 {
		log.Fatal("invalid line: ", line)
	}
	if _, err := st.Exec(tokens[0], tokens[1], tokens[2]); err != nil {
		log.Fatal(err)
	}
}

// END_HL OMIT

func main() {
	db, err := sql.Open("postgres", "postgres://u:p@localhost/d?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(50)
	createTable(db)

	// START_MAIN OMIT
	st, err := db.Prepare("insert into " + table + " values ($1, $2, $3)")
	if err != nil {
		log.Fatal("prepare: ", err)
	}
	defer st.Close()

	s := bufio.NewScanner(os.Stdin)
	wg := sync.WaitGroup{}
	for s.Scan() {
		wg.Add(1)
		go func(line string) {
			handleLine(st, line) // HL
			wg.Done()
		}(s.Text())
	}
	wg.Wait()
	// END_MAIN OMIT
}

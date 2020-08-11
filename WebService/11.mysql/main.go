package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	// "log"
	"io"
)

var db *sql.DB
var err error

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully connected")
	check(err)
}

func friends(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT Name from friends")
	check(err)

	var name string = ""
	var s string = "Received Names: \n"
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`Create table customer (name varchar(45), age int);`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Created table customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer values ("Jacque", 25);`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT name FROM customer`)
	check(err)

	var s string = "Names gotten \n"
	var name string

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Println(w, s)
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jake" where name="Jacque"`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED ROWS: ", n)
}

func delete(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer where name="Jake"`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED ROWS: ", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP table customer`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE: ", n)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Welcome to program")

	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "root:PenguKnight1994!@tcp(localhost:3306)/testdb?charset=utf8")
	check(err)

	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/friends", friends)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}
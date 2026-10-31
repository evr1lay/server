package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3")

var server_port = 8080

var database *sql.DB
var err error

func init_db() {
	database, err = sql.Open("sqlite3", "./forms.db")
	if err != nil {
		panic(err)
	}

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS form (id INTEGER PRIMARY KEY, title TEXT, description TEXT)")
	statement.Exec()
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rows, _ := database.Query("SELECT id, title, description FROM form")
		var id int
		var title, description string

		if rows.Err() != nil {
			panic(rows)
		}

		for rows.Next() {
			rows.Scan(&id, &title, &description)
			fmt.Fprintf(w, "id %d: %s %s\n", id, title, description)
		}
	} else if r.Method == http.MethodPost {
		statement, _ := database.Prepare("INSERT INTO form (title, description) VALUES (?, ?)")
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
			return
		}

		var title, description = r.FormValue("title"), r.FormValue("description")
		statement.Exec(title, description)
		fmt.Printf("%s and %s is added", title, description)
	} else if r.Method == http.MethodDelete {

	}
}

func main() {
	init_db()
	http.HandleFunc("/forms", formHandler)

	fmt.Printf("server is started on http://localhost:%d\n", server_port)

	addr := fmt.Sprintf(":%d", server_port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Please, make sure that this port is free:%d\n", server_port)
	}
}

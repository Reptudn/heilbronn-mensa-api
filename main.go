package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func PrepareDB(db *sql.DB) {

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS menu (id INTEGER PRIMARY KEY, date TEXT, lang TEXT)")
	statement.Exec()

	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS dish (id INTEGER PRIMARY KEY, menu_id INTEGER, name TEXT, type TEXT, FOREIGN KEY(menu_id) REFERENCES menu(id))")
	statement.Exec()

	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS price (id INTEGER PRIMARY KEY, dish_id INTEGER, student TEXT, guest TEXT, worker TEXT, FOREIGN KEY(dish_id) REFERENCES dish(id))")
	statement.Exec()

}

func main() {
	log.Default().Println("Starting server...")

	router := mux.NewRouter()
	db, err := sql.Open("sqlite3", "./menu.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	PrepareDB(db)
	SetupRoutes(router)

	http.ListenAndServe("localhost:4242", router)
}

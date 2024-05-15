package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Price struct {
	Student string `json:"student"`
	Guest   string `json:"guest"`
	Worker  string `json:"worker"`
}

type Dish struct {
	Name  string `json:"name"`
	Price Price  `json:"price"`
	Type  string `json:"type"`
}

type MenuData struct {
	Dishes []Dish `json:"dishes"`
	Lang   string `json:"lang"`
}

type Menu struct {
	MenuData MenuData `json:"menu"`
	Date     string   `json:"date"`
}

func GetMenuByDate(db *sql.DB, date string, location string) *Menu {

	log.Default().Println("Getting menu for", date)

	statement, _ := db.Prepare("SELECT * FROM menu WHERE date=? AND location=?")
	statement.Query(date, location)

	return nil

}

func AddMenu(db *sql.DB, data Menu) bool {
	log.Default().Println("Adding menu")
	return true
}

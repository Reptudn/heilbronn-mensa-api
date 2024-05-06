package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {

	log.Default().Println("Setting up routes...")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("Hello, World!")

	}).Methods("GET")

	r.HandleFunc("/api/menu/{lang}/{date}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		lang := vars["lang"]
		date := vars["date"]

		log.Printf("Received request for menu for %s on %s", lang, date)

		if lang != "en" && lang != "de" {
			http.Error(w, "Invalid language", http.StatusBadRequest)
			return
		}

		_, err := time.Parse("02-01-2006", date)
		if err != nil {
			http.Error(w, "Invalid date", http.StatusBadRequest)
			return
		}

		menu := GetMenuData(date, lang)
		if menu == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		menuJson, err := json.Marshal(menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(menuJson)
	}).Methods("GET")

	log.Default().Println("Routes setup.")

}

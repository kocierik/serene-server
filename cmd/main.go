package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kocierik/SwiftServe/src/db"
	"github.com/kocierik/SwiftServe/src/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/songs", h.GetAllSongs).Methods(http.MethodGet)
	router.HandleFunc("/songs/{id}", h.GetSong).Methods(http.MethodGet)
	router.HandleFunc("/songs", h.AddSong).Methods(http.MethodPost)
	router.HandleFunc("/songs/{id}", h.UpdateSong).Methods(http.MethodPut)
	router.HandleFunc("/songs/{id}", h.DeleteSong).Methods(http.MethodDelete)

  log.Println("API is running...    port 3000")
	http.ListenAndServe(":3000", router)
}

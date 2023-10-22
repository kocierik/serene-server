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
	router.HandleFunc("/songs/{id}", h.GetSongById).Methods(http.MethodGet)
	router.HandleFunc("/songs", h.AddSong).Methods(http.MethodPost)
	router.HandleFunc("/songName/{songName}", h.GetSongByArtistTitle).Methods(http.MethodGet)
	router.HandleFunc("/songs/{id}", h.UpdateSong).Methods(http.MethodPut)
	router.HandleFunc("/songs/{id}", h.DeleteSong).Methods(http.MethodDelete)
	router.HandleFunc("/search", h.GetVideoYt).Methods(http.MethodGet)

	log.Println("API is running...    port 4000")
	http.ListenAndServe(":4000", router)
}

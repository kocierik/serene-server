package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhowden/tag"
	"github.com/kocierik/SwiftServe/src/models"
	"github.com/kocierik/SwiftServe/src/utils"
)

func (h handler) AddSong(w http.ResponseWriter, r *http.Request) {
	mp3File, _, err := r.FormFile("song")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer mp3File.Close()

	mp3Tag, err := tag.ReadFrom(mp3File)
	if err != nil {
		http.Error(w, "Failed to read MP3 tags", http.StatusBadRequest)
		return
	}
	songDuration := utils.GetLengthSong(mp3File)

	song := models.Music{
		Title:    mp3Tag.Title(),
		Artist:   mp3Tag.Artist(),
		Album:    mp3Tag.Album(),
		Duration: songDuration,
		Picture:  mp3Tag.Picture().Data,
	}

	if result := h.DB.Create(&song); result.Error != nil {
		http.Error(w, "Failed to save metadata to the database", http.StatusInternalServerError)
		return
	}

	fmt.Println("DONE: POST ADD SONG")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

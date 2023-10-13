package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kocierik/SwiftServe/src/models"
)

func (h handler) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	var songs []models.Music

	if result := h.DB.Find(&songs); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}

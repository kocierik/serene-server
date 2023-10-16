package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kocierik/SwiftServe/src/models"
)

type Response struct {
	Songs []models.Music `json:"songs"`
}

func (h handler) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	var songs []models.Music

	if result := h.DB.Find(&songs); result.Error != nil {
		fmt.Println(result.Error)
	}

	response := Response{Songs: songs}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response.Songs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kocierik/SwiftServe/src/models"
)

func (h handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedSong models.Music
	json.Unmarshal(body, &updatedSong)

	var song models.Music

	if result := h.DB.First(&song, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	song.Title = updatedSong.Title
	song.Author = updatedSong.Author
	song.Desc = updatedSong.Desc

	h.DB.Save(&song)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

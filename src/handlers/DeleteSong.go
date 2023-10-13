package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kocierik/SwiftServe/src/models"
)

func (h handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var song models.Music

	if result := h.DB.First(&song, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&song)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

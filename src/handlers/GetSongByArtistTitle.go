package handlers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (h handler) GetSongByArtistTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artistTitle := vars["songName"]
	artistTitleTrimmed := strings.ReplaceAll(artistTitle, " ", "")
	// fmt.Println("artistTitle:", artistTitleTrimmed)

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	filePath := "/etc/music/" + artistTitleTrimmed

	songFile, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer songFile.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=song.mp3")
	w.WriteHeader(http.StatusOK)

	http.ServeContent(w, r, "song.mp3", time.Now(), songFile)
}

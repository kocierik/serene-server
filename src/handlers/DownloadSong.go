package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
)

func (h handler) DownloadSong(w http.ResponseWriter, r *http.Request) {
	queryURL := r.URL.Query().Get("query")

	if queryURL == "" {
		http.Error(w, "L'url non e' valido ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	queryURL = "https://www.youtube.com/watch?v=" + queryURL
	cmd := exec.Command("/bin/bash", "uploadSong.sh", queryURL)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Errore", err)
		http.Error(w, "Errore durante l'esecuzione dello script", http.StatusInternalServerError)
		return
	}
	fmt.Println("Download completato con successo.")
	w.Write([]byte("Download completato con successo."))
}

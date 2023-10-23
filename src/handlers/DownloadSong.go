package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
)

func (h handler) DownloadSong(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/bin/bash", "uploadSong.sh", "https://www.youtube.com/watch?v=lZiaYpD9ZrI")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Errore", err)
		http.Error(w, "Errore durante l'esecuzione dello script", http.StatusInternalServerError)
		return
	}
	fmt.Println("Download completato con successo.")
	w.Write([]byte("Download completato con successo."))
}

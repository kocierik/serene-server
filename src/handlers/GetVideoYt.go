package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func (h handler) GetVideoYt(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		http.Error(w, "Errore nel caricamento del file .env: "+err.Error(), http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	query := r.URL.Query().Get("query")

	if query == "" {
		http.Error(w, "La query non può essere vuota", http.StatusBadRequest)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	endpoint := "https://www.googleapis.com/youtube/v3/search?key=" + apiKey + "&part=snippet&q=" + query + "&type=video"

	resp, err := http.Get(endpoint)
	if err != nil {
		http.Error(w, "Errore nella richiesta HTTP a YouTube: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "La richiesta a YouTube ha restituito uno stato non valido: "+resp.Status, http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Errore nella lettura del corpo della risposta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

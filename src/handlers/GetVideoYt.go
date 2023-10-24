package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type VideoDetails struct {
	Items []struct {
		ContentDetails struct {
			Duration string `json:"duration"`
		} `json:"contentDetails"`
	} `json:"items"`
}

type SearchResultItem struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	ID   struct {
		Kind    string `json:"kind"`
		VideoID string `json:"videoId"`
	} `json:"id"`
	Snippet struct {
		PublishedAt string `json:"publishedAt"`
		ChannelID   string `json:"channelId"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Thumbnails  struct {
			Default struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"default"`
			Medium struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"high"`
		} `json:"thumbnails"`
		ChannelTitle         string `json:"channelTitle"`
		LiveBroadcastContent string `json:"liveBroadcastContent"`
		Duration             int    `json:"duration"`
	} `json:"snippet"`
}

type SearchResults struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []SearchResultItem `json:"items"`
}

func parseDurationToSeconds(duration string) (int, error) {
	var minutes, seconds int

	_, err := fmt.Sscanf(duration, "PT%dM%dS", &minutes, &seconds)
	if err != nil {
		return 0, err
	}

	totalSeconds := (minutes * 60) + seconds
	return totalSeconds, nil
}

func getVideoDetails(videoID string, apiKey string) (VideoDetails, error) {
	endpoint := "https://www.googleapis.com/youtube/v3/videos?key=" + apiKey + "&part=contentDetails&id=" + videoID

	resp, err := http.Get(endpoint)
	if err != nil {
		return VideoDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return VideoDetails{}, fmt.Errorf("La richiesta a YouTube ha restituito uno stato non valido: %s", resp.Status)
	}

	var details VideoDetails
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return VideoDetails{}, err
	}

	return details, nil
}

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

	var searchResults SearchResults
	if err := json.NewDecoder(resp.Body).Decode(&searchResults); err != nil {
		http.Error(w, "Errore nella lettura del corpo della risposta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i, item := range searchResults.Items {
		videoID := item.ID.VideoID
		videoDetails, err := getVideoDetails(videoID, apiKey)
		if err != nil {
			http.Error(w, "Errore nell'ottenere i dettagli del video: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if len(videoDetails.Items) > 0 {
			durationInSeconds, err := parseDurationToSeconds(videoDetails.Items[0].ContentDetails.Duration)
			if err != nil {
				http.Error(w, "Errore nella conversione della durata: "+err.Error(), http.StatusInternalServerError)
				return
			}
			searchResults.Items[i].Snippet.Duration = durationInSeconds
		}
	}

	responseJSON, err := json.Marshal(searchResults)
	if err != nil {
		http.Error(w, "Errore nella serializzazione della risposta JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

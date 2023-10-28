package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
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
	if strings.Contains(duration, "H") {
		return 0, nil
	}

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

func (h handler) GetVideoYt(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Errore nel caricamento del file .env: " + err.Error()})
		return
	}

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	query := c.Query("query")

	if query == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "La query non può essere vuota"})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")

	endpoint := "https://www.googleapis.com/youtube/v3/search?key=" + apiKey + "&part=snippet&q=" + query + "&type=video&maxResults=10"

	resp, err := http.Get(endpoint)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Errore nella richiesta HTTP a YouTube: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "La richiesta a YouTube ha restituito uno stato non valido: " + resp.Status})
		return
	}

	var searchResults SearchResults
	if err := json.NewDecoder(resp.Body).Decode(&searchResults); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Errore nella lettura del corpo della risposta: " + err.Error()})
		return
	}

	for i, item := range searchResults.Items {
		videoID := item.ID.VideoID
		videoDetails, err := getVideoDetails(videoID, apiKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Errore nell'ottenere i dettagli del video: " + err.Error()})
			return
		}
		if len(videoDetails.Items) > 0 {
			durationInSeconds, err := parseDurationToSeconds(videoDetails.Items[0].ContentDetails.Duration)
			if durationInSeconds < 600 {
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Errore nella conversione della durata"})
					return
				}
				searchResults.Items[i].Snippet.Duration = durationInSeconds
			}
		}
	}

	c.JSON(http.StatusOK, searchResults)
}

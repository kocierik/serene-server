package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhowden/tag"
	"github.com/gin-gonic/gin"
	"github.com/kocierik/serene-server/src/models"
	"github.com/kocierik/serene-server/src/utils"
)

func (h handler) AddSong(c *gin.Context) {
	mp3File, _, err := c.Request.FormFile("song")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}
	defer mp3File.Close()

	mp3Tag, err := tag.ReadFrom(mp3File)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to read MP3 tags"})
		return
	}
	songDuration := utils.GetLengthSong(mp3File)

	song := models.Music{
		Title:    mp3Tag.Title(),
		Artist:   mp3Tag.Artist(),
		Album:    mp3Tag.Album(),
		Duration: songDuration,
		Picture:  mp3Tag.Picture().Data,
	}

	if result := h.DB.Create(&song); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata to the database"})
		return
	}

	fmt.Println("DONE: POST ADD SONG")
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)
	json.NewEncoder(c.Writer).Encode(song)
}

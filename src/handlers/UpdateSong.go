package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kocierik/SwiftServe/src/models"
)

func (h handler) UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	defer c.Request.Body.Close()
	body, err := io.ReadAll(c.Request.Body)

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

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

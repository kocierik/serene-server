package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kocierik/serene-server/src/models"
)

type Response struct {
	Songs []models.Music `json:"songs"`
}

func (h handler) GetAllSongs(c *gin.Context) {
	var songs []models.Music

	if result := h.DB.Find(&songs); result.Error != nil {
		fmt.Println(result.Error)
	}

	response := Response{Songs: songs}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)

	if err := json.NewEncoder(c.Writer).Encode(response.Songs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

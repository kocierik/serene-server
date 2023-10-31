package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kocierik/serene-server/src/models"
)

func (h handler) DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var song models.Music

	if result := h.DB.First(&song, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&song)
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
	json.NewEncoder(c.Writer).Encode("Deleted")
}

package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if c.Request.Method == http.MethodOptions {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(http.StatusOK)
		return
	}

	idString := fmt.Sprintf("%v", id)
	filePath := "/etc/music/" + idString
	songFile, err := os.Open(filePath)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	defer songFile.Close()

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Content-Disposition", "attachment; filename=song.mp3")
	c.Status(http.StatusOK)

	http.ServeContent(c.Writer, c.Request, "song.mp3", time.Now(), songFile)
}

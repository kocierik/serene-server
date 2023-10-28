package handlers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongByArtistTitle(c *gin.Context) {
	artistTitle := c.Param("songName")
	artistTitleTrimmed := strings.ReplaceAll(artistTitle, " ", "")

	if c.Request.Method == http.MethodOptions {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(http.StatusOK)
		return
	}

	filePath := "/etc/music/" + artistTitleTrimmed

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

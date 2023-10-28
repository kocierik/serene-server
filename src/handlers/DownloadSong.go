package handlers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func (h handler) DownloadSong(c *gin.Context) {
	queryURL := c.Query("query")

	if queryURL == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "L'url non e' valido"})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)

	queryURL = "https://www.youtube.com/watch?v=" + queryURL
	cmd := exec.Command("/bin/bash", "uploadSong.sh", queryURL)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Errore", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Errore durante l'esecuzione dello script"})
		return
	}
	fmt.Println("Download completato con successo.")
	c.Writer.Write([]byte("Download completato con successo."))
}

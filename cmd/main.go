package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kocierik/serene-server/src/db"
	"github.com/kocierik/serene-server/src/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := gin.Default()

	router.GET("/songs", func(c *gin.Context) {
		h.GetAllSongs(c)
	})
	router.GET("/songs/:id", func(c *gin.Context) {
		h.GetSongById(c)
	})
	router.POST("/songs", func(c *gin.Context) {
		h.AddSong(c)
	})
	router.GET("/songName/:songName", func(c *gin.Context) {
		h.GetSongByArtistTitle(c)
	})
	router.PUT("/songs/:id", func(c *gin.Context) {
		h.UpdateSong(c)
	})
	router.DELETE("/songs/:id", func(c *gin.Context) {
		h.DeleteSong(c)
	})
	router.GET("/search", func(c *gin.Context) {
		h.GetVideoYt(c)
	})
	router.GET("/downloadSong", func(c *gin.Context) {
		h.DownloadSong(c)
	})

	log.Println("API is running...    port 4000")
	http.ListenAndServe(":4000", router)
}

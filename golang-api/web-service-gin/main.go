package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	album{ID: "1", Title: "Blue Train", Artist: "", Price: 56.99},
	album{ID: "2", Title: "Jeru", Artist: "", Price: 17.99},
	album{ID: "3", Title: "Music", Artist: "Blue", Price: 10.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

//here is the handler to return a specific item (post and get are working)

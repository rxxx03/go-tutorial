package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type album struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}

var albums = []album{
	{ID: uuid.MustParse("59c8b985-655d-408e-93d9-3ec46324fb6e"), Title: "Loveless", Artist: "My Bloody Valentine", Price: 43.99},
	{ID: uuid.MustParse("a1a93f6b-44fd-4841-8d7d-34cd3faae7b7"), Title: "Beatles", Artist: "Rubber Soul", Price: 75.00},
	{ID: uuid.MustParse("0c9560e4-5d59-433e-b966-57aa03a8f295"), Title: "Lynyrd Skynyrd", Artist: "Freebird", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid UUID"})
		return
	}

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum.ID)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

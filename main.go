package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.PUT("/albums/:id", putAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//not working
func putAlbum(c *gin.Context) {
	id := c.Param("id")
	var newAlbuminfo album

	var album album
	for _, i := range albums {
		if i.ID != id {
			if err := c.BindJSON(&album); err != nil {
				return
			}
		}
	}
	fmt.Print(album)
	album.Title = newAlbuminfo.Title
	album.Artist = newAlbuminfo.Artist
	album.Price = newAlbuminfo.Price

	c.IndentedJSON(http.StatusAccepted, albums)
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	var newAlbums []album
	index := 0
	for _, i := range albums {
		if i.ID != id {
			newAlbums = append(newAlbums, i)
			index++
		}
	}
	c.IndentedJSON(http.StatusCreated, newAlbums)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

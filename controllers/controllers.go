package controllers

import (
	"fmt"
	"net/http"

	"github.com/Shreyas-Prabhu/MONGOCRUDGO/helpers"
	"github.com/Shreyas-Prabhu/MONGOCRUDGO/models"
	"github.com/gin-gonic/gin"
)

func InsertController(c *gin.Context) {
	var movie models.Movie
	c.Bind(&movie)
	str := helpers.InsertMovie(movie)
	fmt.Println(str)
	c.JSON(http.StatusOK, gin.H{
		"Inserted Movie": movie,
	})
}

func GetOneController(c *gin.Context) {
	id := c.Param("id")
	movie := helpers.GetOneMovie(id)
	c.JSON(http.StatusOK, gin.H{
		"Got Movie": movie,
	})
}

func GetAllController(c *gin.Context) {
	movie := helpers.GetAllMovies()
	c.JSON(http.StatusOK, gin.H{
		"Got Movies": movie,
	})
}

func UpdateMovieController(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	c.Bind(&movie)
	movie1 := helpers.UpdateMovie(id, movie)
	c.JSON(http.StatusOK, gin.H{
		"Updated Movie": movie1,
	})
}

func DeleteOneController(c *gin.Context) {
	id := c.Param("id")
	str := helpers.DeleteOneMovie(id)
	c.JSON(http.StatusOK, gin.H{
		"Message": str,
	})
}

func DeleteAllController(c *gin.Context) {
	str := helpers.DeleteAll()
	c.JSON(http.StatusOK, gin.H{
		"Message": str,
	})
}

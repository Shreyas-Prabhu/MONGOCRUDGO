package main

import (
	"fmt"
	"net/http"

	"github.com/Shreyas-Prabhu/MONGOCRUDGO/controllers"
	"github.com/Shreyas-Prabhu/MONGOCRUDGO/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitConfig()
}

func main() {
	fmt.Println("Welcome to mongodb CRUD...3010 port used")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to MONGO CRUD USING GIN",
		})
	})
	r.GET("/movies", controllers.GetAllController)
	r.GET("/movie/:id", controllers.GetOneController)
	r.POST("/movie", controllers.InsertController)
	r.PUT("/movie/:id", controllers.UpdateMovieController)
	r.DELETE("/movie/:id", controllers.DeleteOneController)
	r.DELETE("/movies", controllers.DeleteAllController)
	r.Run()

}

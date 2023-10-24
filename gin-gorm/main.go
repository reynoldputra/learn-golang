package main

import (
	"gin-gorm/controllers"
	"gin-gorm/models"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
  Author string `json:"author"`
}

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", controllers.GetAllBooks)
	r.GET("/:id", controllers.FindBook)
	r.POST("/", controllers.CreateBook)
	r.PATCH("/:id", controllers.UpdateBook)
	r.DELETE("/:id", controllers.DeleteBook)

	r.Run()
}

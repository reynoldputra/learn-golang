package controllers

import (
	"gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks (c *gin.Context) {
  var books []models.Book

  models.DB.Find(&books)

  c.JSON(http.StatusOK, gin.H{"data" : books})
}

func FindBook (c *gin.Context) {
  var book models.Book
  
  err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error" : err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data" : book})
}

func CreateBook (c *gin.Context) {
  var input models.CreateBook
  err := c.ShouldBindJSON(&input)

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
    return
  }

  book := models.Book{
    Title: input.Title,
    Author: input.Author,
  }

  models.DB.Create(&book)

  c.JSON(http.StatusOK, gin.H{"data" : book})
}


func UpdateBook (c *gin.Context) {
  var book models.Book
  
  err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error" : "Record not found"})
    return
  }

  var input models.UpdateBook
  err = c.ShouldBindJSON(&input)

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
    return
  }

  models.DB.Model(&book).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data" : book})
}

func DeleteBook (c *gin.Context) {
  var book models.Book
  
  err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error" : "Record not found"})
    return
  }

  models.DB.Delete(&book)

  c.JSON(http.StatusOK, gin.H{"data" : "Delete success"})
}

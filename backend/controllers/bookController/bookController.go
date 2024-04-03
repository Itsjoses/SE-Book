package bookcontroller

import (
	"fmt"
	"net/http"

	"github.com/Itsjoses/sebook-be/database"
	"github.com/Itsjoses/sebook-be/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Preload("Genre").Find(&books)
	// Send the result as a response
	c.JSON(200, books)
}

func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := database.DB.Preload("Genre").Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book Not Found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func SearchBook(c *gin.Context) {
	var search struct {
		Search string
	}

	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	fmt.Print(search)

	var books []models.Book
	if err := database.DB.Where("title LIKE ?", "%"+search.Search+"%").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

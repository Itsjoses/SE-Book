package routes

import (
	bookcontroller "github.com/Itsjoses/sebook-be/controllers/bookController"
	seedercontroller "github.com/Itsjoses/sebook-be/controllers/seederController"
	usercontroller "github.com/Itsjoses/sebook-be/controllers/userController"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/signup", usercontroller.Signup)

	r.GET("/genreseed", seedercontroller.GenreSeeder)
	r.GET("/bookbseed", seedercontroller.BookSeeder)

	r.GET("/getbooks", bookcontroller.GetBooks)
	r.GET("/getbook/:id", bookcontroller.GetBook)
	r.POST("/searchbook", bookcontroller.SearchBook)
}

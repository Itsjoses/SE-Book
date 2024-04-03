package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type Genre struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type Book struct {
	gorm.Model
	Title   string
	Cover   int
	Price   int
	GenreID uint
	Genre   Genre `gorm:"foreignkey:GenreID"`
}

package models

import (
	"day_2/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Quantity  int    `json:"quantity"`
}

func init() {
	config.InitDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func UpdateBook(b *Book, id uint) *Book {
	db.Where("id = ?", id).Updates(b)
	return b
}

func DeleteBook(id uint) {
	db.Delete(&Book{}, id)
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(id uint) *Book {
	var book Book
	db.Where("id =?", id).First(&book)
	return &book
}

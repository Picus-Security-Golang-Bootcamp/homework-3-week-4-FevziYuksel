package book

import (
	"encoding/json"
	"fmt"
	"homework3/author"
	"io/ioutil"
	"os"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name      string        `json:"Name"`
	Pages     uint64        `json:"Pages"`
	Stocks    uint64        `json:"Stocks"`
	Price     uint64        `json:"Price"`
	StockCode string        `json:"StockCode"`
	ISBN      string        `json:"ISBN"`
	Author_ID uint64        `json:"Author_ID"`
	Author    author.Author `gorm:"foreignKey:Author_ID"  json:"Author"`
}

type BookSlice struct {
	Books []Book `json:"Books"`
}

func (b *BookSlice) ReadJSON(fileAdress string) {
	file, err := os.Open(fileAdress)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	jsonFile, _ := ioutil.ReadAll(file)
	err = json.Unmarshal([]byte(jsonFile), b)
	if err != nil {
		fmt.Println(err)
	}
}

func (b *BookSlice) ConvertBook() []Book {
	var books []Book
	books = append(books, b.Books...)
	return books
}

func (b *BookSlice) ExtractAuthor() []author.Author {
	var authors []author.Author
	for _, eachAuthor := range b.Books {
		authors = append(authors, eachAuthor.Author)
	}
	return authors
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s", b.ID, b.Name)
}
func (Book) TableName() string {
	return "books"
}

func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleting...", b.Name)
	return nil
}

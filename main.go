package main

import (
	"fmt"
	"homework3/author"
	book "homework3/book"
	postgres "homework3/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	var bookSlice1 book.BookSlice
	bookSlice1.ReadJSON("BookList.json")
	books := bookSlice1.ConvertBook()
	// authors := bookSlice1.ExtractAuthor()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()

	bookRepo := book.NewBookRepository(db)
	bookRepo.Migrations()

	// authorRepo.InsertSampleData(authors)
	bookRepo.InsertSampleData(books)
	// fmt.Println(bookRepo.FindAll())
	fmt.Println(bookRepo.FindByID(0))
}

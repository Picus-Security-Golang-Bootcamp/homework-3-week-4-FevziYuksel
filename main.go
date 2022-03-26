package main

import (
	postgres "homework3/db"
	"homework3/domain/models"
	"homework3/domain/repos"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//JSON to struct methods
	var bookSlice1 models.BookSlice
	bookSlice1.ReadJSON("BookList.json")
	// authors := bookSlice1.ExtractAuthor()
	books := bookSlice1.ConvertBook()

	// DB connection and create repositories
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	authorRepo := repos.NewAuthorRepository(db)
	authorRepo.Migration()
	bookRepo := repos.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertSampleData(books)

	// fmt.Println(bookRepo.FindAll())  		//find all books
	// fmt.Println(bookRepo.GetByID(1)) 			//get books by ids
	// fmt.Println(bookRepo.FindByName("C")) 		//find books by names
	// fmt.Println(bookRepo.FindByNameWithRawSQL("C"))
	// fmt.Println(bookRepo.DeleteById(1))
	// fmt.Println(bookRepo.GetBooksWithAuthorInformation())
	// fmt.Println(authorRepo.GetAuthorWithBookInformation())

}

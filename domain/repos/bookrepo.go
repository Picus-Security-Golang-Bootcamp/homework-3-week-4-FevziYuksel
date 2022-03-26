package repos

import (
	"errors"
	"homework3/domain/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&models.Book{})

}
func (b *BookRepository) Create(book models.Book) error {
	result := b.db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) InsertSampleData(books []models.Book) {

	for _, book := range books {
		newBook := models.Book{
			Name:      book.Name,
			Pages:     book.Pages,
			Stocks:    book.Stocks,
			Price:     book.Price,
			StockCode: book.StockCode,
			ISBN:      book.ISBN,
			Author_ID: book.Author_ID,
			Author:    book.Author,
		}
		b.db.Where(models.Book{Name: newBook.Name}).FirstOrCreate(&newBook)
	}
}

func (b *BookRepository) FindAll() []models.Book {
	var books []models.Book
	b.db.Find(&books)
	return books
}

func (b *BookRepository) GetByID(id int) (*models.Book, error) {
	var book models.Book
	result := b.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &book, nil
}

func (b *BookRepository) FindByName(name string) []models.Book {
	var book []models.Book
	b.db.Where("name LIKE ? ", "%"+name+"%").Find(&book)

	return book
}

func (b *BookRepository) FindByNameWithRawSQL(name string) []models.Book {
	var books []models.Book
	b.db.Raw("SELECT * FROM books WHERE name LIKE ?", "%"+name+"%").Scan(&books)

	return books
}

func (b *BookRepository) Delete(book models.Book) error {
	result := b.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BookRepository) DeleteById(id uint64) error {
	result := b.db.Delete(&models.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) GetBooksWithAuthorInformation() ([]models.Book, error) {
	var books []models.Book
	result := b.db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

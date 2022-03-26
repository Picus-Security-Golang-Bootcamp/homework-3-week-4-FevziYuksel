package book

import (
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})

}
func (b *BookRepository) Create(book Book) error {
	result := b.db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) InsertSampleData(books []Book) {

	for _, book := range books {
		newBook := Book{
			Book_ID:   book.Book_ID,
			Name:      book.Name,
			Pages:     book.Pages,
			Stocks:    book.Stocks,
			Price:     book.Price,
			StockCode: book.StockCode,
			ISBN:      book.ISBN,
			Author_ID: book.Author_ID,
			Author:    book.Author, // Bu da tablo oluşturuyor
		}
		b.db.Where(Book{Book_ID: newBook.Book_ID}).FirstOrCreate(&newBook)
	}
}

func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)
	return books
}

func (b *BookRepository) FindByID(ID uint64) []Book {
	var books []Book
	b.db.Where(`"Book_ID" = ?`, ID).Order("Id desc,name").Find(&books)
	b.db.Where(&Book{Book_ID: ID}).First(&books)
	b.db.Where(map[string]interface{}{"Book_ID": ID, "code": "01"}).First(&books)
	return books
}

/*
func (b *BookRepository) FindByCountryCodeOrCityCode(code string) []Book {
	var books []Book
	b.db.Where(`"CountryCode = ?"`, code).Or("code = ?", code).Find(&books)

	return books
}

func (b *BookRepository) FindByName(name string) []Book {
	var books []Book
	b.db.Where("name LIKE ? ", "%"+name+"%").Find(&books)

	return books
}

//*
func (b *BookRepository) FindByNameWithRawSQL(name string) []Book {
	var books []Book
	b.db.Raw("SELECT * FROM City WHERE name LIKE ?", "%"+name+"%").Scan(&books)

	return books
}

//*
func (b *BookRepository) GetByID(id int) (*Book, error) {
	var books Book
	result := b.db.First(&books, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &books, nil
}



func (b *BookRepository) Update(book Book) error {
	result := b.db.Save(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) Delete(book Book) error {
	result := b.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BookRepository) DeleteById(id int) error {
	result := b.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
*/

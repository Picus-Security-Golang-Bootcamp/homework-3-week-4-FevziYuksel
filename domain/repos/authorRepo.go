package repos

import (
	"homework3/domain/models"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}
func (a *AuthorRepository) Migration() {
	a.db.AutoMigrate(&models.Author{})
}

func (a *AuthorRepository) InsertSampleData(authors []models.Author) {
	authorList := []models.Author{}
	for _, author := range authors {
		newAuthors := models.Author{
			Author_Name: author.Author_Name,
		}
		authorList = append(authorList, newAuthors)
	}
	for _, eachAuthor := range authorList {
		a.db.Create(&eachAuthor)

	}
}

func (a *AuthorRepository) GetAuthorWithBookInformation() ([]models.Author, error) {
	var authors []models.Author
	result := a.db.Preload("Book").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

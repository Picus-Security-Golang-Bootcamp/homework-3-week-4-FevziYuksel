package author

import (
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}
func (a *AuthorRepository) Migration() {
	a.db.AutoMigrate(&Author{})
}

func (a *AuthorRepository) InsertSampleData(authors []Author) {
	authorList := []Author{}
	for _, author := range authors {
		newAuthors := Author{
			Author_ID:   author.Author_ID,
			Author_Name: author.Author_Name,
		}
		authorList = append(authorList, newAuthors)
	}
	// Bu da tablo oluşturuyor
	for _, eachAuthor := range authorList {
		a.db.Create(&eachAuthor)
		// a.db.Where(Author{Author_ID: eachAuthor.Author_ID}).Attrs(Author{
		// 	Author_Name: eachAuthor.Author_Name,
		// }).FirstOrCreate(&eachAuthor)
	}
	// for _, eachAuthor := range authors {
	// 	// a.db.Where(Author{Author_ID: eachAuthor.Author_ID}).FirstOrCreate(&eachAuthor)
	// 	a.db.Create(&eachAuthor)
	// }
}

// func (a *AuthorRepository) GetAllCountriesWithCityInformation() ([]Author, error) {
// 	var authors []Author
// 	result := c.db.Preload("Cities").Find(&authors)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return countries, nil
// }

// func (c *CountryRepository) GetCountryWithName(name string) (*Country, error) {
// 	var country *Country
// 	result := c.db.
// 		Where(Country{Name: name}).
// 		Attrs(Country{Code: "NULL", Name: "NULL"}).
// 		FirstOrInit(&country) // Eğer sorgu sonucunda veri bulunursa Attrs kısmında yazılanlar ignore edilir.

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return country, nil
// }

// func (c *CountryRepository) GetAllCountriesWithoutCityInformation() ([]Country, error) {
// 	var countries []Country
// 	result := c.db.Find(&countries)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return countries, nil
// }

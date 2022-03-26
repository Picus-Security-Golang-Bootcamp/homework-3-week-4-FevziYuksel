package author

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Author_ID   uint64 ` json:"Author_ID" ` //
	Author_Name string `json:"Author_Name"`
}
type AuthorSlice []Author

func (a *Author) TableName() string {
	return "authors"
}
func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s", a.Author_ID, a.Author_Name)
}
func (a *Author) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Author (%s) deleting...", a.Author_Name)
	return nil
}

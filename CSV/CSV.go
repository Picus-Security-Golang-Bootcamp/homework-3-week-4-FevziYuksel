package CSV

import (
	"encoding/csv"
	"os"
	"strconv"

	author "homework3/author"
	book "homework3/book"
)

// func main() {
// 	_, _, err := ReadCSVAuthor("models.csv")
// 	fmt.Print(err)
// 	// fmt.Printf("%v\n,%v\n,%v\n", s1[0], s2[0], err)

// }
func ReadCSVAuthor(filename string) (author.AuthorSlice, book.BookSlice, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	listAuthor := author.AuthorSlice{}
	for _, line := range records[1:] {
		listAuthor = append(listAuthor, author.Author{
			Author_ID:   stringToUInt(line[7]),
			Author_Name: line[8],
		})
	}
	listBook := book.BookSlice{}
	for i, line := range records[1:] {
		// fmt.Println(line[0])
		// fmt.Printf("%v\n", stringToUInt(line[2]))
		// fmt.Printf("%T\n", line[2])
		listBook = append(listBook, book.Book{
			ID:        stringToUInt(line[0]),
			Name:      line[1],
			Pages:     stringToUInt(line[2]),
			Stocks:    stringToUInt(line[3]),
			Price:     stringToUInt(line[4]),
			StockCode: line[5],
			ISBN:      line[6],
			Author_ID: stringToUInt(line[7]),
			Author:    listAuthor[i],
		})
	}
	// for _, l := range listBook {
	// 	fmt.Println(l.ToString())
	// }
	return listAuthor, listBook, nil
}
func stringToUInt(str string) uint64 {
	num, _ := strconv.ParseUint(str, 10, 64)
	return num
}

package main

import (
	"hw1/book"
	bookmap "hw1/bookstorage/map"
	bookslice "hw1/bookstorage/slice"
	"hw1/library"
)

func GetGenerator(start_id int) func() int {
	id := start_id
	return func() int {
		id++
		return id
	}
}
func main() {
	lib := &library.Library{Ids: make(map[string][]int)}
	lib.SetBookstorage(bookmap.GetEmptyStorage())
	lib.SetGenerator(GetGenerator(1))
	var bookList []book.Book = []book.Book{
		{Name: "Book1", Author: "Author1", Date: "2023-01-01"},
		{Name: "Book2", Author: "Author2", Date: "2023-02-01"},
		{Name: "Book3", Author: "Author3", Date: "2023-03-01"},
		{Name: "Book4", Author: "Author4", Date: "2023-04-01"},
		{Name: "Book5", Author: "Author5", Date: "2023-05-01"},
		{Name: "Book1", Author: "Author1", Date: "2023-01-01"},
	}
	for i := range bookList {
		lib.AddBook(&bookList[i])
	}
	B := lib.GetBookByName("Book1")
	println(B.Name, B.Author, B.Date)
	B = lib.GetBookByName("Book1")
	println(B.Name, B.Author, B.Date)
	B = lib.GetBookByName("Book1")
	if B == nil {
		println("No more Book1 available")
	}
	lib.SetGenerator(GetGenerator(100))
	B = lib.GetBookByName("Book2")
	println(B.Name, B.Author, B.Date)
	lib.SetBookstorage(bookslice.GetEmptyStorage())
	lib.AddBook(&book.Book{Name: "NewBook", Author: "NewAuthor", Date: "2024-01-01"})
	lib.AddBook(&book.Book{Name: "AnotherBook", Author: "AnotherAuthor", Date: "2024-02-01"})
	lib.AddBook(&book.Book{Name: "Another Love", Author: "Tom Odell", Date: "2024-03-01"})
	B = lib.GetBookByName("NewBook")
	println(B.Name, B.Author, B.Date)
	B = lib.GetBookByName("AnotherBook")
	println(B.Name, B.Author, B.Date)

}

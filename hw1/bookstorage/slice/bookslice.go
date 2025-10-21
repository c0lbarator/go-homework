package bookslice

import (
	"hw1/book"
	"hw1/library"
)

type BookstorageSlice struct {
	Books []*book.Book
}

func (b *BookstorageSlice) AddBook(book *book.Book) {
	b.Books = append(b.Books, book)
}

func (b *BookstorageSlice) GetBookByID(id int) *book.Book {
	for i, book := range b.Books {
		if book.Id == id {
			// Remove the book from the slice
			b.Books = append(b.Books[:i], b.Books[i+1:]...)
			return book
		}
	}
	return nil
}
func (b *BookstorageSlice) GetEmptyStorage() library.Storage {
	return &BookstorageSlice{
		Books: []*book.Book{},
	}
}
func GetEmptyStorage() library.Storage {
	return &BookstorageSlice{
		Books: []*book.Book{},
	}
}
func (b *BookstorageSlice) GetBooks() []*book.Book {
	return b.Books
}

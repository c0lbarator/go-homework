package bookmap

import (
	"hw1/book"
	"hw1/library"
)

type BookstorageMap struct {
	Books map[int]*book.Book
}

func (b *BookstorageMap) AddBook(book *book.Book) {
	b.Books[book.Id] = book
}

func (b *BookstorageMap) GetBookByID(id int) *book.Book {
	book := b.Books[id]
	delete(b.Books, id)
	return book
}

func (b *BookstorageMap) GetEmptyStorage() library.Storage {
	return &BookstorageMap{
		Books: make(map[int]*book.Book),
	}
}
func GetEmptyStorage() library.Storage {
	return &BookstorageMap{
		Books: make(map[int]*book.Book),
	}
}
func (b *BookstorageMap) GetBooks() []*book.Book {
	books := make([]*book.Book, 0, len(b.Books))
	for _, book := range b.Books {
		books = append(books, book)
	}
	return books
}

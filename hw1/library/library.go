package library

import (
	"hw1/book"
)

type Storage interface {
	AddBook(b *book.Book)
	GetBookByID(id int) *book.Book
	GetEmptyStorage() Storage
	GetBooks() []*book.Book
}
type Library struct {
	Bookstorage Storage
	Ids         map[string][]int
	Generator   func() int
}

func (l *Library) AddBook(b *book.Book) {
	id := l.Generator()
	b.Id = id
	l.Bookstorage.AddBook(b)
	if l.Ids[b.Name] == nil {
		l.Ids[b.Name] = make([]int, 0)
	}
	l.Ids[b.Name] = append(l.Ids[b.Name], id)
}
func (l *Library) GetBookByName(name string) *book.Book {
	ids, ok := l.Ids[name]
	if ok && len(ids) > 0 {
		book := l.Bookstorage.GetBookByID(ids[0])
		l.Ids[name] = ids[1:]
		if len(l.Ids[name]) == 0 {
			delete(l.Ids, name)
		}
		return book
	}
	return nil
}
func (l *Library) SetGenerator(f func() int) {
	l.Generator = f
	newIds := make(map[string][]int)
	newBookstorage := l.Bookstorage.GetEmptyStorage()
	for name, ids := range l.Ids {
		for _, id := range ids {
			book := l.Bookstorage.GetBookByID(id)
			newId := f()
			book.Id = newId
			if newIds[name] == nil {
				newIds[name] = make([]int, 0)
			}
			newIds[name] = append(newIds[name], newId)
			newBookstorage.AddBook(book)
		}
	}
	l.Bookstorage = newBookstorage
	l.Ids = newIds
}
func (l *Library) SetBookstorage(b Storage) {
	l.Bookstorage = b
	l.Ids = make(map[string][]int)
	for _, book := range b.GetBooks() {
		if l.Ids[book.Name] == nil {
			l.Ids[book.Name] = make([]int, 0)
		}
		l.Ids[book.Name] = append(l.Ids[book.Name], book.Id)
	}
}

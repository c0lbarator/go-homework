package main

type Book struct {
	// name, id, author, date
	Id     int
	Name   string
	Author string
	Date   string
}
type Bookstorage struct {
	Books map[int]*Book
}
type Library struct {
	Bookstorage Bookstorage
	Ids         map[string][]int
	Generator   func() int
}

func (l *Library) AddBook(b *Book) {
	id := l.Generator()
	l.Bookstorage.Books[id] = b
	if l.Ids[b.Name] == nil {
		l.Ids[b.Name] = make([]int, 0)
	}
	l.Ids[b.Name] = append(l.Ids[b.Name], id)
}
func (l *Library) GetBookByName(name string) *Book {
	ids, ok := l.Ids[name]
	if ok && len(ids) > 0 {
		book := l.Bookstorage.Books[ids[0]]
		delete(l.Bookstorage.Books, ids[0])
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
	newBookstorage := Bookstorage{
		Books: make(map[int]*Book),
	}
	newIds := make(map[string][]int)
	for _, book := range l.Bookstorage.Books {
		newId := f()
		book.Id = newId
		newBookstorage.Books[newId] = book
		if newIds[book.Name] == nil {
			newIds[book.Name] = make([]int, 0)
		}
		newIds[book.Name] = append(newIds[book.Name], newId)
	}
	l.Bookstorage = newBookstorage
	l.Ids = newIds
}
func (l *Library) SetBookstorage(b Bookstorage) {
	l.Bookstorage = b
	for id, book := range b.Books {
		if l.Ids[book.Name] == nil {
			l.Ids[book.Name] = make([]int, 0)
		}
		l.Ids[book.Name] = append(l.Ids[book.Name], id)
	}
}
func GetGenerator(start_id int) func() int {
	id := start_id
	return func() int {
		id++
		return id
	}
}
func main() {
	lib := &Library{Ids: make(map[string][]int)}
	lib.SetGenerator(GetGenerator(1))
	lib.SetBookstorage(Bookstorage{
		Books: make(map[int]*Book),
	})
	var books []Book = []Book{
		{Name: "Book1", Author: "Author1", Date: "2023-01-01"},
		{Name: "Book2", Author: "Author2", Date: "2023-02-01"},
		{Name: "Book3", Author: "Author3", Date: "2023-03-01"},
		{Name: "Book4", Author: "Author4", Date: "2023-04-01"},
		{Name: "Book5", Author: "Author5", Date: "2023-05-01"},
		{Name: "Book1", Author: "Author1", Date: "2023-01-01"},
	}
	for i := range books {
		lib.AddBook(&books[i])
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
	lib.SetBookstorage(Bookstorage{
		Books: make(map[int]*Book),
	})
	lib.AddBook(&Book{Name: "NewBook", Author: "NewAuthor", Date: "2024-01-01"})
	lib.AddBook(&Book{Name: "AnotherBook", Author: "AnotherAuthor", Date: "2024-02-01"})
	lib.AddBook(&Book{Name: "Another Love", Author: "Tom Odell", Date: "2024-03-01"})
	B = lib.GetBookByName("NewBook")
	println(B.Name, B.Author, B.Date)
	B = lib.GetBookByName("AnotherBook")
	println(B.Name, B.Author, B.Date)

}

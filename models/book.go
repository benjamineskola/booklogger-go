package models

type Book struct {
	ID    uint
	Title string
}

func NewBook(title string) Book {
	return Book{ //nolint:exhaustivestruct
		Title: title,
	}
}

func (book Book) String() string {
	return book.Title
}

func (book Book) FullTitle() string {
	return book.Title
}

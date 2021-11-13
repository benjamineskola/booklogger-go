package models

import "time"

type Book struct {
	ID                uint
	Title             string
	Subtitle          string
	FirstAuthor       string
	FirstAuthorRole   string
	AdditionalAuthors []string
	Slug              string

	FirstPublished uint
	Language       string
	Series         string
	SeriesOrder    uint

	EditionPublished uint
	Publisher        string
	EditionFormat    uint
	EditionNumber    uint
	PageCount        uint
	GoodreadsID      string
	GoogleBooksID    string
	Isbn             string
	Asin             string
	EditionLanguage  string
	EditionTitle     string
	EditionSubtitle  string

	AcquiredDate  time.Time
	AlienatedDate time.Time
	WasBorrowed   bool
	BorrowedFrom  string

	OwnedBy      string
	ImageURL     string
	PublisherURL string
	WantToRead   string
	Tags         []string
	Review       string
	Rating       uint

	HasEbookEdition   bool
	EbookAcquiredDate time.Time
	EbookAsin         string
	EbookIsbn         string

	Editions []*Book `gorm:"many2many:library_book_editions"`

	ParentEditionID *uint
	ParentEdition   *Book

	Private bool

	ModifiedDate time.Time `gorm:"autoUpdateTime"`
	CreatedDate  time.Time `gorm:"autoCreateTime"`
}

func (Book) TableName() string {
	return "library_book"
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
	if book.Subtitle != "" {
		return book.Title + ": " + book.Subtitle
	}

	return book.Title
}

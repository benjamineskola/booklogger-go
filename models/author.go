package models

import (
	"fmt"
	"strings"
	"time"
)

type Author struct {
	ID                 uint
	Surname            string
	Forenames          string
	PreferredForenames string
	Slug               string

	FirstAuthoredBooks      []Book `gorm:"foreignKey:FirstAuthorID"`
	AdditionalAuthoredBooks []Book `gorm:"many2many:library_bookauthor"`

	Gender          uint
	POC             bool
	SurnameFirst    bool
	PrimaryLanguage string

	PrimaryIdentityID *uint
	PrimaryIdentity   *Author

	ModifiedDate time.Time `gorm:"autoUpdateTime"`
	CreatedDate  time.Time `gorm:"autoCreateTime"`
}

func (Author) TableName() string {
	return "library_author" // notest
}

func NewAuthor(name string) *Author {
	words := strings.Split(name, " ")
	surname, words := words[len(words)-1], words[:len(words)-1]

	prefixes := map[string]bool{
		"le": true, "de": true, "von": true, "der": true, "van": true,
	}

	for i := len(words); i > 0; i-- {
		word := words[len(words)-i]
		_, inPrefixes := prefixes[word]

		if inPrefixes {
			surname, words = words[len(words)-1]+" "+surname, words[:len(words)-1]
		}
	}

	return &Author{ //nolint:exhaustivestruct
		Surname:   surname,
		Forenames: strings.ReplaceAll(strings.Join(words, " "), ". ", "."),
	}
}

func (author Author) String() string {
	return author.DisplayName()
}

func (author Author) DisplayName() string {
	if author.Forenames == "" {
		return author.Surname
	}

	var forenames string
	if author.PreferredForenames != "" {
		forenames = author.PreferredForenames
	} else {
		forenames = author.Forenames
	}

	if author.SurnameFirst {
		return fmt.Sprintf("%s %s", author.Surname, forenames)
	}

	return fmt.Sprintf("%s %s", forenames, author.Surname)
}

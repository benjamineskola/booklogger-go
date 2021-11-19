package models_test

import (
	"booklogger/models"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()

	englishTitle := "Capital"
	germanTitle := "Das Kapital"
	englishSubtitle := "Critique of Political Economy"
	germanSubtitle := "Kritik der politischen Ökonomie"
	englishTitleFull := englishTitle + ": " + englishSubtitle
	germanTitleFull := germanTitle + ": " + germanSubtitle

	tests := []struct {
		name             string
		book             models.Book
		expected         string
		expectedOriginal string
	}{
		{"simple title", models.Book{Title: englishTitle}, englishTitle, ""},
		{
			"title and subtitle",
			models.Book{Title: englishTitle, Subtitle: englishSubtitle},
			englishTitleFull, "",
		},
		{
			"translation",
			models.Book{
				Title:           germanTitle,
				Subtitle:        germanSubtitle,
				EditionTitle:    englishTitle,
				EditionSubtitle: englishSubtitle,
			},
			englishTitleFull,
			germanTitleFull,
		},
		{
			"translation without subtitle",
			models.Book{Title: germanTitle, EditionTitle: englishTitle},
			englishTitle,
			germanTitle,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.book.String()
			if actual != test.expected {
				t.Fatalf("expected: %s; got: %s", test.expected, actual)
			}

			if test.expectedOriginal != "" {
				actual = test.book.OriginalTitle()
				if actual != test.expectedOriginal {
					t.Fatalf("expected: %s; got: %s", test.expectedOriginal, actual)
				}
			}
		})
	}
}

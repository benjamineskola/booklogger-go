package models_test

import (
	"booklogger/models"
	"testing"
)

func TestBook(t *testing.T) { //nolint:funlen
	t.Parallel()

	englishTitle := "Capital"
	germanTitle := "Das Kapital"
	englishSubtitle := "Critique of Political Economy"
	germanSubtitle := "Kritik der politischen Ökonomie"
	englishTitleFull := englishTitle + ": " + englishSubtitle
	germanTitleFull := germanTitle + ": " + germanSubtitle

	t.Run("simple book with title", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook(englishTitle)
		if book.FullTitle() != englishTitle {
			t.Fatal("title not set")
		}
	})
	t.Run("book with title and subtitle", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook(englishTitle)
		book.Subtitle = englishSubtitle
		if book.FullTitle() != englishTitleFull {
			t.Fatal("full title not set")
		}
	})
	t.Run("String method is same as Fulltitle", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook(englishTitle)
		if book.String() != englishTitle {
			t.Fatal("title not set")
		}
		book.Subtitle = englishSubtitle
		if book.FullTitle() != englishTitleFull {
			t.Fatal("full title not set")
		}
	})

	t.Run("If the edition has a different title, Fulltitle returns it", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook(germanTitle)
		book.Subtitle = germanSubtitle
		book.EditionTitle = englishTitle
		book.EditionSubtitle = englishSubtitle

		if book.FullTitle() != englishTitleFull {
			t.Fatal("full title should be this edition's, not the original title")
		}

		if book.String() != englishTitleFull {
			t.Fatal("string representation should be this edition's, not the original title")
		}

		if book.OriginalTitle() != germanTitleFull {
			t.Fatal("original title should display book's full original title")
		}
	})

	t.Run("Fulltitle works without subtitles too", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook(germanTitle)
		book.EditionTitle = englishTitle
		if book.FullTitle() != englishTitle {
			t.Fatal("full title should be this edition's, not the original title")
		}
		if book.OriginalTitle() != germanTitle {
			t.Fatal("original title should display")
		}
	})
}

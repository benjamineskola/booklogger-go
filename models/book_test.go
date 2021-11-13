package models_test

import (
	"booklogger/models"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()

	t.Run("simple book with title", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook("The Communist Manifesto")
		if book.FullTitle() != "The Communist Manifesto" {
			t.Fatal("title not set")
		}
	})
	t.Run("book with title and subtitle", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook("Capital")
		book.Subtitle = "Critique of Political Economy"
		if book.FullTitle() != "Capital: Critique of Political Economy" {
			t.Fatal("full title not set")
		}
	})
	t.Run("String method is same as Fulltitle", func(t *testing.T) {
		t.Parallel()
		book := models.NewBook("Capital")
		if book.String() != "Capital" {
			t.Fatal("title not set")
		}
		book.Subtitle = "Critique of Political Economy"
		if book.FullTitle() != "Capital: Critique of Political Economy" {
			t.Fatal("full title not set")
		}
	})
}

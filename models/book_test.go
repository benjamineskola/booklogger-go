package models_test

import (
	"booklogger/models"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()

	book := models.NewBook("The Communist Manifesto")
	if book.FullTitle() != "The Communist Manifesto" {
		t.Fatal("title not set")
	}
}
